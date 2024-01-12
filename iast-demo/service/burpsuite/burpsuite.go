package burpsuite

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	googleuuid "github.com/google/uuid"
	"iast-demo/common"
	"iast-demo/entity"
	"iast-demo/entity/config"
	"iast-demo/http"
	"iast-demo/service"
	"iast-demo/util"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

var (
	burpFlag      string
	uuid          string
	taskTargetKey string
	packetKey     string
	taskKey       string
	resultKey     string
	sysType       string
)

func init() {
	uuid = googleuuid.New().String()
	burpFlag = common.GenRedisKey("flag", "burpsuite", uuid)
	packetKey = "packet:burpsuite:"
	taskKey = "task:burpsuite:"
	taskTargetKey = taskKey + "target"
	resultKey = "result:burpsuite:"

	sysType = runtime.GOOS
}
func Run(exeCmd string) {
	var cmd *exec.Cmd
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()

	switch sysType {
	case "windows":
		cmd = exec.Command("cmd.exe", "/k", exeCmd)
	default:
		cmd = exec.Command("/bin/sh", "-c", exeCmd)
	}
	stdoutPipe, err := cmd.StdoutPipe()
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		rdb.Del(ctx, burpFlag)
		log.Printf("start burp rest api failed: %s", err)
	}
	if config.CFG.Burpsuite.ShowLog {
		outReader := bufio.NewReader(stdoutPipe)
		errReader := bufio.NewReader(stderrPipe)
		go common.PrintReader(outReader)
		go common.PrintReader(errReader)
	}
	err = cmd.Wait()
	if err != nil {
		rdb.Del(ctx, burpFlag)
		log.Printf("wait burp rest api failed: %s", err)
	}
}

func SendRequestThroughProxy(flow *entity.Flow) (bool, error) {
	httpData := &entity.HttpData{
		Api:        flow.RawUrl,
		HttpMethod: flow.Method,
		UrlData:    flow.UrlParameter,
		Header:     flow.Header,
	}
	if flow.Header != nil {
		httpData.Header = flow.Header
	}
	if flow.BodyParameter != "" {
		httpData.BodyData = flow.BodyParameter
	}
	proxy := util.GetHttpProxy()
	_, _, err := http.DoRequest(httpData, proxy)
	if err != nil {
		return false, err
	}
	return true, nil
}

func StatusOk() bool {
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	keys, _, err := rdb.Scan(ctx, 0, packetKey+"*", 0).Result()
	if err != nil {
		log.Fatal(err)
	}
	// 没有流量
	if len(keys) == 0 {
		return false
	}
	version, err := GetVersion()
	// 未启动
	if version == nil || version.BurpVersion == "" {
		result, _ := rdb.Exists(ctx, burpFlag).Result()
		if result == 0 {
			log.Println("burpsuite not started")
			log.Println("start burpsuite")
			rdb.Set(ctx, burpFlag, uuid, -1)
			Run(config.CFG.Burpsuite.ExeCmd)
		}
		return false
	}
	return true
}

func SubmitTask() {
	if !StatusOk() {
		return
	}
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	//已存在扫描任务
	result, err := rdb.Exists(ctx, taskTargetKey).Result()
	if result != 0 {
		log.Println("burpsuite scan task running")
		return
	}

	count := config.CFG.Burpsuite.SingleScanCount
	keys, _, err := rdb.Scan(ctx, 0, packetKey+"*", 0).Result()
	//直接按配置的数量pop当前用户流量
	if len(keys) > 0 {
		key := keys[0]
		resultList, err := rdb.RPopCount(ctx, key, count).Result()
		if err != nil {
			log.Println(err)
		}
		targetSet := mapset.NewSet()
		for _, v := range resultList {
			flow := &entity.Flow{}
			json.Unmarshal([]byte(v), flow)
			url := flow.Url
			if !targetSet.Contains(url) {
				targetSet.Add(url)
				rdb.LPush(ctx, taskTargetKey+":"+flow.Username, url)
				targetInScope, _ := GetTargetInScope(url)
				if !targetInScope.InScope {
					_, err := SetTarget2Scope(url)
					if err != nil {
						log.Fatalf("set burpsuite target to scope failed: %s", err)
					}
				}
			}
			_, err = SendRequestThroughProxy(flow)
			log.Printf("send request through burpsuite proxy【%s %s】", key, flow.RawUrl)
			if err != nil {
				log.Printf("send request failed: %s", err)
				continue
			}
			// 还是跟原来的流量一样
			burpTaskData := &entity.BurpsuiteTaskData{
				TaskId: flow.StandardMd5,
				Flow:   flow,
			}
			marshal, _ := json.Marshal(burpTaskData)
			rdb.LPush(ctx, taskKey+flow.RequestHost+":"+flow.Username, marshal)
		}
		for target := range targetSet.Iterator().C {
			_, err := StartActiveScan(target.(string))
			if err != nil {
				log.Fatalf("burpsuite scan failed: %s", err)
			}
			log.Println("start burpsuite active scan", target)
		}
	}
}

func GetTaskResult() {
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	keys, _, err := rdb.Scan(ctx, 0, taskKey+"*", 0).Result()
	if err != nil {
		log.Fatalf("%s", err)
	}
	// 没有任务
	if len(keys) == 0 {
		return
	}
	percent, err := GetScanPercent()
	if err != nil {
		log.Println(err)
	}
	log.Printf("burpsuite scan status: %d", percent)
	if percent == 100 {
		targetKeys, _, err := rdb.Scan(ctx, 0, taskTargetKey+"*", 0).Result()
		if err != nil {
			log.Fatalf("%s", err)
		}
		for _, key := range targetKeys {
			result, _ := rdb.RPopCount(ctx, key, 200).Result()
			username := strings.Split(key, ":")[3]
			for _, url := range result {
				issue, err := GetScanIssues(url)
				if err != nil {
					log.Printf("burpsuite get issues failed: %s", err)
				}
				burpResultData := &entity.BurpsuiteResultData{
					BurpResult: &issue.Issues,
					Flow:       nil,
					RawRequest: "",
					Username:   username,
				}
				marshal, _ := json.Marshal(burpResultData)
				rdb.LPush(ctx, resultKey+issue.Issues[0].Host, marshal)
			}
		}
		_, err = StopBurp()
		if err != nil {
			log.Fatalf("stop burp failed:%s", err)
		}
		rdb.Del(ctx, burpFlag)
		rdb.Del(ctx, taskTargetKey)
		for _, key := range keys {
			rdb.RPopCount(ctx, key, 200)
		}
	}
}

func ShowTaskResult(severities ...string) {
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	result, err := rdb.Keys(ctx, resultKey+"*").Result()
	if len(result) != 0 {
		sqlmapResult := "burpsuite_result.txt"
		file, err := os.OpenFile(sqlmapResult, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()
		if err != nil && os.IsNotExist(err) {
			file, _ = os.Create(sqlmapResult)
		}
		resultWriter := bufio.NewWriter(file)
		iter := rdb.Scan(ctx, 0, resultKey+"*", 0).Iterator()
		for iter.Next(ctx) {
			key := iter.Val()
			resultList, err := rdb.RPopCount(ctx, key, 200).Result()
			if err != nil {
				log.Println(err)
			}
			for _, v := range resultList {
				burpResultData := &entity.BurpsuiteResultData{}
				json.Unmarshal([]byte(v), burpResultData)
				issues := burpResultData.BurpResult
				for _, iss := range *issues {
					var showResult bool
					for _, severity := range severities {
						if strings.ToLower(severity) == "all" || strings.ToLower(severity) == strings.ToLower(iss.Severity) {
							showResult = true
						}
					}
					if len(severities) == 0 || showResult {
						lineStart := strings.Repeat("*", 10) + "详情" + strings.Repeat("*", 10)
						data1 := fmt.Sprintf("漏洞：%s\n链接：%s\n风险：%s\n可信：%s", iss.IssueName, iss.Url, iss.Severity, iss.Confidence)
						data2 := fmt.Sprintf("协议：%s\n主机：%s\n端口：%d", iss.Protocol, iss.Host, iss.Port)

						fmt.Println(lineStart)
						//fmt.Println(data1)
						switch iss.Severity {
						case "High":
							//红
							fmt.Printf("\033[1;31m%s\033[0m\n", data1)
						case "Medium":
							//黄
							fmt.Printf("\033[1;33m%s\033[0m\n", data1)
						case "Low":
							//蓝
							fmt.Printf("\033[1;34m%s\033[0m\n", data1)
						case "Information":
							//灰
							fmt.Printf("\033[1;30m%s\033[0m\n", data1)
						}

						fmt.Println(data2)
						fmt.Fprintln(resultWriter, lineStart)
						fmt.Fprintln(resultWriter, data1)

						if len(iss.HttpMessages) != 0 {
							for _, msg := range iss.HttpMessages {
								data3 := fmt.Sprintf("方法：%s\nURL：%s", msg.Method, msg.Url)
								fmt.Println(data3)
								fmt.Fprintln(resultWriter, data3)
								if len(msg.Parameters) != 0 {
									linePara := fmt.Sprintf("参数：%+v", msg.Parameters)
									fmt.Println(linePara)
									fmt.Fprintln(resultWriter, linePara)
								}
								// reqDecoded, _ := base64.StdEncoding.DecodeString(msg.Request)
								// fmt.Println("Request", string(reqDecoded))
								// respDecoded, _ := base64.StdEncoding.DecodeString(msg.Response)
								// fmt.Println("Response", string(respDecoded))
							}
						}
						lineDetail := fmt.Sprintf("细节：%s", iss.IssueDetail)
						fmt.Println(lineDetail)
						fmt.Fprintln(resultWriter, lineDetail)
						lineEnd := strings.Repeat("*", 25)
						fmt.Fprintln(resultWriter, lineEnd+"\n")
						fmt.Println(lineEnd)
					}
				}
			}
		}
		resultWriter.Flush()
	}

}
