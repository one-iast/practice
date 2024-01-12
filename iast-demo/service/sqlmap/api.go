package sqlmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"iast-demo/entity"
	"iast-demo/http"
)

var baseApi = "http://127.0.0.1:8775"

// task api
func NewTask() (string, error) {
	httpData := &entity.HttpData{Api: baseApi + "/task/new", HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return "", err
	}
	result := &entity.NewSqlmapTaskResp{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.Taskid, nil
}

func DeleteScanTask(taskId string) (bool, error) {
	httpData := &entity.HttpData{Api: baseApi + fmt.Sprintf("/task/%s/delete", taskId), HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}

	resultData, err := checkFalse(body, err)
	if err != nil {
		return false, err
	}
	result := &entity.Success{}
	err = json.Unmarshal(resultData, &result)
	if err != nil {
		return false, err
	}
	return result.Success, nil
}

// task api end

// scan api

func StartScan(taskId string, data *entity.SqlmapTaskFlow) (int, error) {
	headerMap := map[string]string{}
	headerMap["Content-Type"] = "application/json"
	httpData := &entity.HttpData{Api: baseApi + fmt.Sprintf("/scan/%s/start", taskId), HttpMethod: "POST", BodyData: data, Header: headerMap}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return 0, err
	}
	resultData, err := checkFalse(body, err)
	if err != nil {
		return 0, err
	}
	result := &entity.StartSqlmapTaskResp{}
	err = json.Unmarshal(resultData, &result)
	if err != nil {
		return 0, err
	}
	if !result.Success.Success {
		return 0, errors.New("unknown err ")
	}
	return result.EngineId, nil
}

func GetScanLog(taskId string) (*entity.SqlmapLogResp, error) {
	httpData := &entity.HttpData{Api: baseApi + fmt.Sprintf("/scan/%s/log", taskId), HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}

	resultData, err := checkFalse(body, err)
	if err != nil {
		return nil, err
	}
	result := &entity.SqlmapLogResp{}
	err = json.Unmarshal(resultData, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetScanStatus(taskId string) (string, error) {
	httpData := &entity.HttpData{Api: baseApi + fmt.Sprintf("/scan/%s/status", taskId), HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return "", err
	}

	resultData, err := checkFalse(body, err)
	if err != nil {
		return "", err
	}
	result := &entity.SqlmapTaskEndResp{}
	err = json.Unmarshal(resultData, &result)
	if err != nil {
		return "", err
	}
	return result.Status, nil
}

func GetScanData(taskId string) (*entity.SqlmapResultResp, error) {
	httpData := &entity.HttpData{Api: baseApi + fmt.Sprintf("/scan/%s/data", taskId), HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	resultData, err := checkFalse(body, err)
	if err != nil {
		return nil, err
	}
	result := &entity.SqlmapResultResp{}
	err = json.Unmarshal(resultData, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// scan api end

// admin api

func GetScanList() (*entity.SqlmapRunningTask, error) {
	httpData := &entity.HttpData{Api: baseApi + "/admin/list", HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.SqlmapRunningTask{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// admin api end

func CheckStatus() (bool, error) {
	httpData := &entity.HttpData{Api: baseApi, HttpMethod: "GET"}
	_, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}
	return true, nil
}
