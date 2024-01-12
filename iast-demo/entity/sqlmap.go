package entity

type Success struct {
	Success bool `json:"success"`
}
type SqlmapTaskFalseResp struct {
	Success
	Message string `json:"message"`
}
type NewSqlmapTaskResp struct {
	Success
	Taskid string `json:"taskid"`
}
type StartSqlmapTaskResp struct {
	Success
	EngineId int `json:"engineid"`
}
type SqlmapTaskEndResp struct {
	Success
	Status     string `json:"status"`
	ReturnCode any    `json:"returncode"`
}

type SqlmapLog struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type SqlmapLogResp struct {
	Success
	Log []SqlmapLog `json:"log"`
}

type SqlmapDataType0 struct {
	Url   string `json:"url"`
	Query string `json:"query"`
	Data  string `json:"data"`
}
type SqlmapDataType1 struct {
	Place     string `json:"place"`
	Parameter string `json:"parameter"`
	Ptype     int    `json:"ptype"`
	Prefix    string `json:"prefix"`
	Suffix    string `json:"suffix"`
	Clause    []int  `json:"clause"`
	Notes     []any  `json:"notes"`
	Data      any    `json:"data"`
	Conf      struct {
		TextOnly  any `json:"textOnly"`
		Titles    any `json:"titles"`
		Code      any `json:"code"`
		String    any `json:"string"`
		NotString any `json:"notString"`
		Regexp    any `json:"regexp"`
		Optimize  any `json:"optimize"`
	} `json:"conf"`
	Dbms        string   `json:"dbms"`
	DbmsVersion []string `json:"dbms_version"`
	Os          any      `json:"os"`
}
type SqlmapDataType2 struct {
	Status int    `json:"status"`
	Type   int    `json:"type"`
	Value  string `json:"value"`
}
type SqlmapPayloadDetail struct {
	Title           string `json:"title"`
	Payload         string `json:"payload"`
	Where           int    `json:"where"`
	Vector          []any  `json:"vector"`
	Comment         string `json:"comment"`
	TemplatePayload any    `json:"templatePayload"`
	MatchRatio      any    `json:"matchRatio"`
	TrueCode        any    `json:"trueCode"`
	FalseCode       any    `json:"falseCode"`
}
type SqlmapResultRespData struct {
	Status int `json:"status"`
	Type   int `json:"type"`
	Value  any `json:"value"`
}
type SqlmapResultResp struct {
	Success bool                   `json:"success"`
	Data    []SqlmapResultRespData `json:"data"`
	Error   []any                  `json:"error"`
}

// 如果还有其它的type，需要额外加，现在只有：0、1、2
type SqlmapResult struct {
	UrlInfo      *SqlmapDataType0 `json:"url_info,omitempty"`
	InjectDetail *SqlmapDataType1 `json:"inject_detail,omitempty"`
	// InjectDetail的payload单独拎出来
	PayloadList []*SqlmapPayloadDetail `json:"payload_list,omitempty"`
	HostInfo    *SqlmapDataType2       `json:"host_info,omitempty"`
}

type SqlmapRunningTask struct {
	Success  bool              `json:"success"`
	Tasks    map[string]string `json:"tasks"`
	TasksNum int               `json:"tasks_num"`
}

type SqlmapTaskFlow struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	Data    string `json:"data,omitempty" `
	Headers string `json:"headers,omitempty"`
	Cookie  string `json:"cookie,omitempty"`
	Agent   string `json:"agent,omitempty"`
	Host    string `json:"host,omitempty"`
	Referer string `json:"referer,omitempty"`
}

type SqlmapPacketData struct {
	SqlmapTaskFlow *SqlmapTaskFlow `json:"sqlmap_task_flow"`
	Flow           *Flow           `json:"flow"`
}

type SqlmapTaskData struct {
	TaskId string `json:"task_id"`
	Flow   *Flow  `json:"flow"`
}
type SqlmapResultData struct {
	SqlmapResult *SqlmapResult `json:"sqlmap_result"`
	Flow         *Flow         `json:"flow"`
	RawRequest   string        `json:"raw_request"`
}
