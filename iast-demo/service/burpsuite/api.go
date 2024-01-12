package burpsuite

import (
	"encoding/json"
	"fmt"
	"iast-demo/entity"
	"iast-demo/http"
)

var (
	baseApi = "http://localhost:8090/burp"
)

func GetVersion() (*entity.BurpVersion, error) {
	httpData := &entity.HttpData{Api: baseApi + "/versions", HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.BurpVersion{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func StopBurp() (bool, error) {
	httpData := &entity.HttpData{Api: baseApi + "/stop", HttpMethod: "GET"}
	_, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetProxyHistory() (*entity.BurpProxyHistory, error) {
	httpData := &entity.HttpData{Api: baseApi + "/proxy/history", HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.BurpProxyHistory{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetConfig() (*entity.BurpConfig, error) {
	httpData := &entity.HttpData{Api: baseApi + "/configuration", HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.BurpConfig{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SetConfig(conf *entity.BurpConfig) (bool, error) {
	httpData := &entity.HttpData{Api: baseApi + "/configuration", HttpMethod: "PUT", BodyData: conf}
	_, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Query whether a specific URL is within the current Suite-wide scope. Returns true if an url is in scope.
func GetTargetInScope(target string) (*entity.BurpScope, error) {
	httpData := &entity.HttpData{Api: baseApi + "/target/scope", HttpMethod: "GET", UrlData: "url=" + target}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.BurpScope{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Includes the specified URL in the Suite-wide scope.
func SetTarget2Scope(target string) (bool, error) {
	httpData := &entity.HttpData{Api: baseApi + "/target/scope", HttpMethod: "PUT", UrlData: "url=" + target}
	_, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Excludes the specified Url from the Suite-wide scope.
func DeleteTargetFromScope(target string) (bool, error) {
	httpData := &entity.HttpData{Api: baseApi + "/target/scope", HttpMethod: "DELETE", UrlData: "url=" + target}
	_, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}
	return true, nil
}

// StartActiveScan 会对该目标sitemap中所有流量进行扫描（即使已经扫过了）
// Scans through Burp Sitemap and sends all HTTP requests with url starting with baseUrl to Burp Scanner for active scan.
// 需要删除site map，才不会再扫描，ui有删除的功能，但是接口没有给删除的，所以通过接口删除不了site map
func StartActiveScan(target string) (bool, error) {
	httpData := &entity.HttpData{Api: baseApi + "/scanner/scans/active", HttpMethod: "POST", UrlData: "baseUrl=" + target}
	_, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteScanQueueMap Deletes the scan queue map from memory, not from Burp suite UI.
func DeleteScanQueueMap() (bool, error) {
	httpData := &entity.HttpData{Api: baseApi + "/scanner/scans/active", HttpMethod: "DELETE"}
	_, _, err := http.DoRequest(httpData)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Please note that the scanner status percentages returned by Burp v1.7 and v2.x don’t have the same granularity.
func GetScanStatus() (*entity.BurpScanStatus, error) {
	httpData := &entity.HttpData{Api: baseApi + "/scanner/status", HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.BurpScanStatus{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Returns status details of items in the scan queue. For percentage, use /scanner/status instead.
func GetScanDetail() (*entity.BurpScanDetail, error) {
	httpData := &entity.HttpData{Api: baseApi + "/scanner/status/details", HttpMethod: "GET"}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.BurpScanDetail{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetScanIssues(urlPrefix string) (*entity.BurpScanIssue, error) {
	httpData := &entity.HttpData{Api: baseApi + "/scanner/issues", HttpMethod: "GET", UrlData: "urlPrefix=" + urlPrefix}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	result := &entity.BurpScanIssue{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 需要将xml解析成struct，有点难度，可以用"github.com/beevik/etree"
// GetScanIssues已经可以覆盖所有的结果了
func GetScanReport(urlPrefix string) ([]byte, error) {
	//urlPrefix, reportType, issueSeverity, issueConfidence
	//HTML | XML
	reportType := "XML"
	//All, High, Medium, Low and Information，可以多个逗号分开
	issueSeverity := "All"
	// All, Certain, Firm and Tentative，可以多个逗号分开
	issueConfidence := "All"
	httpData := &entity.HttpData{Api: baseApi + "/report", HttpMethod: "GET", UrlData: fmt.Sprintf("urlPrefix=%s&reportType=%s&issueSeverity=%s&issueConfidence=%s", urlPrefix, reportType, issueSeverity, issueConfidence)}
	body, _, err := http.DoRequest(httpData)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetScanPercent() (int, error) {
	status, err := GetScanStatus()
	if err != nil {
		return 0, err
	}
	detail, err := GetScanDetail()
	//fmt.Printf("%+v\n", detail)
	if err != nil {
		return 0, err
	}
	if len(detail.Scans) != 0 {
		return status.ScanPercentage, nil
	}
	return 0, nil
}
