package entity

type HttpData struct {
	Api        string            `json:"api"`
	HttpMethod string            `json:"httpMethod"`
	Header     map[string]string `json:"header"`
	UrlData    string            `json:"urlData"`
	BodyData   any               `json:"bodyData"`
}

type ReqData struct {
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	Data      any    `json:"data"`
}
