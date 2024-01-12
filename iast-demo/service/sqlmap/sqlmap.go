package sqlmap

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"iast-demo/common"
	"iast-demo/entity"
	"iast-demo/entity/config"
	"iast-demo/service"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Run(sqlmapApiPath string) {
	cmd := exec.Command(config.CFG.Stage.PythonPath, sqlmapApiPath, "-s")
	stdoutPipe, err := cmd.StdoutPipe()
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Printf("start sqlmapapi failed: %s", err)
	}
	if config.CFG.Sqlmap.ShowLog {
		outReader := bufio.NewReader(stdoutPipe)
		errReader := bufio.NewReader(stderrPipe)
		go common.PrintReader(outReader)
		go common.PrintReader(errReader)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("wait sqlmapapi failed: %s", err)
	}
}

// GetTaskCount 获取特定任务的数量
// default fetch running task count
func GetTaskCount(rType ...int) (int, error) {
	var runType int
	if len(rType) != 0 {
		runType = rType[0]
	}
	taskType := "running"
	switch runType {
	case 1:
		taskType = "terminated"
	case 2:
		taskType = "not running"
	}
	list, err := GetScanList()
	if err != nil {
		return 0, err
	}
	var runningTask []string
	for _, v := range list.Tasks {
		if v == taskType {
			runningTask = append(runningTask, v)
		}
	}
	return len(runningTask), nil
}
func checkFalse(body []byte, err error) ([]byte, error) {
	var resultData map[string]any
	json.Unmarshal(body, &resultData)
	success := resultData["success"].(bool)
	if success {
		return body, nil
	} else {
		result := &entity.SqlmapTaskFalseResp{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(result.Message)
	}
}

func SubmitTask() {
	if !StatusOk() {
		return
	}
	var (
		count int
		err   error
	)
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	iter := rdb.Scan(ctx, 0, "packet:sqlmap*", 0).Iterator()
	for iter.Next(ctx) {
		count, err = GetTaskCount()
		if err != nil {
			log.Println(err)
		}
		i := config.CFG.Sqlmap.Thread - count
		if count == config.CFG.Sqlmap.Thread || i <= 0 {
			return
		}
		key := iter.Val()
		resultList, err := rdb.RPopCount(ctx, key, i).Result()
		if err != nil {
			log.Println(err)
		}
		for _, v := range resultList {
			flow := &entity.SqlmapPacketData{}
			json.Unmarshal([]byte(v), flow)
			taskId, err := NewTask()
			log.Printf("submit task【%s %s %s】", key, taskId, flow.Flow.RawUrl)
			if err != nil {
				log.Fatalf("submit task failed: %s", err)
			}
			_, err = StartScan(taskId, flow.SqlmapTaskFlow)
			if err != nil {
				log.Fatalf("start task failed: %s", err)
			}
			sqlmapTaskData := &entity.SqlmapTaskData{
				TaskId: taskId,
				Flow:   flow.Flow,
			}
			marshal, _ := json.Marshal(sqlmapTaskData)
			rdb.LPush(ctx, "task:sqlmap:"+taskId, marshal)
		}
	}
}

func GetTaskResult() {
	var (
		count int
		err   error
	)
	count, err = GetTaskCount()
	if err != nil {
		log.Println(err)
	}
	i := config.CFG.Sqlmap.Thread - count
	if count == config.CFG.Sqlmap.Thread || i <= 0 {
		return
	}
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	iter := rdb.Scan(ctx, 0, "task:sqlmap*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		split := strings.Split(key, ":")
		taskId := split[len(split)-1]
		status, err := GetScanStatus(taskId)
		if err != nil {
			log.Printf("get task status failed: %s", err)
			if err.Error() == "Invalid task ID" {
				log.Printf("remove task: %s\n", taskId)
				DeleteScanTask(taskId)
				rdb.Del(ctx, key)
			}
		}
		//已结束
		if status == "terminated" {
			v, _ := rdb.RPop(ctx, key).Result()
			taskData := &entity.SqlmapTaskData{}
			json.Unmarshal([]byte(v), taskData)
			log.Printf("parse task result【%s %s】", key, taskData.Flow.RawUrl)
			scanData, err := GetScanData(taskId)
			if err != nil {
				log.Fatal(err)
			}
			_, err = DeleteScanTask(taskId)
			if err != nil {
				log.Fatalf("delete task failed: %s", err)
			}
			result := ParseSqlmapResult(scanData)
			if result != nil && result.UrlInfo != nil {
				flow := taskData.Flow
				var rawRequest []string
				firstLine := fmt.Sprintf("%s %s %s", flow.Method, flow.PrettyUrl, flow.HttpVersion)
				rawRequest = append(rawRequest, firstLine)
				//headers
				rawRequest = append(rawRequest, common.Map2StringLine(flow.Header))
				//body
				if flow.BodyParameter != "" {
					rawRequest = append(rawRequest, "")
					rawRequest = append(rawRequest, flow.BodyParameter)
				}
				sqlmapResultData := &entity.SqlmapResultData{
					SqlmapResult: result,
					Flow:         flow,
					RawRequest:   strings.Join(rawRequest, "\n"),
				}

				marshal, _ := json.Marshal(sqlmapResultData)
				rdb.LPush(ctx, "result:sqlmap:"+flow.RequestHost, marshal)
			} else {
				fmt.Printf("\033[1;32m%s\033[0m【%s %s】\n", "[x] No SQL injection:", taskData.Flow.Method, taskData.Flow.RawUrl)
			}
		}
	}
}

func ShowTaskResult() {
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	result, err := rdb.Keys(ctx, "result:sqlmap*").Result()
	if len(result) != 0 {
		sqlmapResult := "sqlmap_result.txt"
		file, err := os.OpenFile(sqlmapResult, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()
		if err != nil && os.IsNotExist(err) {
			file, _ = os.Create(sqlmapResult)
		}
		resultWriter := bufio.NewWriter(file)
		iter := rdb.Scan(ctx, 0, "result:sqlmap*", 0).Iterator()
		for iter.Next(ctx) {
			key := iter.Val()
			resultList, err := rdb.RPopCount(ctx, key, 200).Result()
			if err != nil {
				log.Println(err)
			}
			for _, v := range resultList {
				sqlmapResultData := &entity.SqlmapResultData{}
				json.Unmarshal([]byte(v), sqlmapResultData)
				log.Printf("show task result【%s %s】", key, sqlmapResultData.Flow.RawUrl)

				fmt.Printf("\033[1;31m%s\033[0m%s\n", "[!] SQL injection found: ", v)

				//fmt.Printf("%+v\n", sqlmapResultData)
				fmt.Fprintln(resultWriter, v)
			}
		}
		resultWriter.Flush()
	}
}

func ParseSqlmapResult(scanData *entity.SqlmapResultResp) *entity.SqlmapResult {
	sqlmapResult := &entity.SqlmapResult{}
	for _, data := range scanData.Data {
		dataValue := data.Value
		switch data.Type {
		case 0:
			// url数据
			sqlmapUrlDetail := &entity.SqlmapDataType0{}
			mapstructure.Decode(dataValue, sqlmapUrlDetail)
			sqlmapResult.UrlInfo = sqlmapUrlDetail
		case 1:
			// 注入点信息
			for _, data := range dataValue.([]any) {
				sqlmapDataValue1 := &entity.SqlmapDataType1{}
				mapstructure.Decode(data, sqlmapDataValue1)
				sqlmapResult.InjectDetail = sqlmapDataValue1

				//详细的注入点信息
				var payloadDetails []*entity.SqlmapPayloadDetail
				for _, v := range sqlmapDataValue1.Data.(map[string]any) {
					sqlmapPayloadDetail := &entity.SqlmapPayloadDetail{}
					mapstructure.Decode(v, sqlmapPayloadDetail)
					payloadDetails = append(payloadDetails, sqlmapPayloadDetail)
				}

				sqlmapResult.PayloadList = payloadDetails
			}
		case 2:
			// 主机信息
			sqlmapDataValue2 := &entity.SqlmapDataType2{}
			mapstructure.Decode(data, sqlmapDataValue2)
			sqlmapResult.HostInfo = sqlmapDataValue2
		}
	}
	return sqlmapResult
}

func StatusOk() bool {
	_, err := CheckStatus()
	if err != nil {
		Run(config.CFG.Sqlmap.Path)
		return false
	}
	return true
}
