package entity

import "time"

type BurpVersion struct {
	BurpVersion      string `json:"burpVersion"`
	ExtensionVersion string `json:"extensionVersion"`
}

type BurpProxyHistory struct {
	Messages []struct {
		Comment string `json:"comment"`
		Cookies []struct {
			Domain     string    `json:"domain"`
			Expiration time.Time `json:"expiration"`
			Name       string    `json:"name"`
			Path       string    `json:"path"`
			Value      string    `json:"value"`
		} `json:"cookies"`
		Highlight  string `json:"highlight"`
		Host       string `json:"host"`
		Method     string `json:"method"`
		Parameters []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"parameters"`
		Port            int      `json:"port"`
		Protocol        string   `json:"protocol"`
		Request         string   `json:"request"`
		Response        string   `json:"response"`
		ResponseHeaders []string `json:"responseHeaders"`
		StatusCode      int      `json:"statusCode"`
		Url             string   `json:"url"`
		//Url             struct {
		//	Authority string `json:"authority"`
		//	Content   struct {
		//	} `json:"content"`
		//	DefaultPort int    `json:"defaultPort"`
		//	File        string `json:"file"`
		//	Host        string `json:"host"`
		//	Path        string `json:"path"`
		//	Port        int    `json:"port"`
		//	Protocol    string `json:"protocol"`
		//	Query       string `json:"query"`
		//	Ref         string `json:"ref"`
		//	UserInfo    string `json:"userInfo"`
		//} `json:"url"`
	} `json:"messages"`
}

type BurpScope struct {
	Url     string `json:"url"`
	InScope bool   `json:"inScope"`
}
type BurpScanStatus struct {
	ScanPercentage int `json:"scanPercentage"`
}

type BurpScanDetail struct {
	Scans []struct {
		Url    string `json:"url"`
		Status string `json:"status"`
	} `json:"scans"`
}

type BurpScanIssue1 struct {
	Issues []struct {
		Url                   string      `json:"url"`
		IssueName             string      `json:"issueName"`
		IssueType             int         `json:"issueType"`
		Severity              string      `json:"severity"`
		Confidence            string      `json:"confidence"`
		IssueBackground       string      `json:"issueBackground"`
		IssueDetail           *string     `json:"issueDetail"`
		RemediationBackground string      `json:"remediationBackground"`
		RemediationDetail     interface{} `json:"remediationDetail"`
		HttpMessages          []struct {
			Host            string   `json:"host"`
			Port            int      `json:"port"`
			Protocol        string   `json:"protocol"`
			Url             string   `json:"url"`
			StatusCode      int      `json:"statusCode"`
			Request         string   `json:"request"`
			Response        string   `json:"response"`
			Comment         string   `json:"comment"`
			Highlight       string   `json:"highlight"`
			Method          string   `json:"method"`
			ResponseHeaders []string `json:"responseHeaders"`
			Cookies         []struct {
				Domain     interface{} `json:"domain"`
				Expiration interface{} `json:"expiration"`
				Name       string      `json:"name"`
				Path       string      `json:"path"`
				Value      string      `json:"value"`
			} `json:"cookies"`
			Parameters []interface{} `json:"parameters"`
		} `json:"httpMessages"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Protocol string `json:"protocol"`
	} `json:"issues"`
}

type Issues struct {
	Confidence   string `json:"confidence"`
	Host         string `json:"host"`
	HttpMessages []struct {
		Comment string `json:"comment"`
		Cookies []struct {
			Domain     string    `json:"domain"`
			Expiration time.Time `json:"expiration"`
			Name       string    `json:"name"`
			Path       string    `json:"path"`
			Value      string    `json:"value"`
		} `json:"cookies"`
		Highlight  string `json:"highlight"`
		Host       string `json:"host"`
		Method     string `json:"method"`
		Parameters []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"parameters"`
		Port            int      `json:"port"`
		Protocol        string   `json:"protocol"`
		Request         string   `json:"request"`
		Response        string   `json:"response"`
		ResponseHeaders []string `json:"responseHeaders"`
		StatusCode      int      `json:"statusCode"`
		Url             string   `json:"url"`
	} `json:"httpMessages"`
	IssueBackground       string `json:"issueBackground"`
	IssueDetail           string `json:"issueDetail"`
	IssueName             string `json:"issueName"`
	IssueType             int    `json:"issueType"`
	Port                  int    `json:"port"`
	Protocol              string `json:"protocol"`
	RemediationBackground string `json:"remediationBackground"`
	RemediationDetail     string `json:"remediationDetail"`
	Severity              string `json:"severity"`
	Url                   string `json:"url"`
}
type BurpScanIssue struct {
	Issues []Issues `json:"issues"`
}

type BurpsuiteTaskData struct {
	TaskId string `json:"task_id"`
	Flow   *Flow  `json:"flow"`
}
type BurpsuiteResultData struct {
	Username   string    `json:"username"`
	BurpResult *[]Issues `json:"burp_result"`
	Flow       *Flow     `json:"flow"`
	RawRequest string    `json:"raw_request"`
}

type BurpScanData struct {
	Username string
	Url      string
}
