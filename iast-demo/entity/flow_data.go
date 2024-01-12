package entity

type Flow struct {
	Username      string            `json:"username"`
	Url           string            `json:"url"`        //协议、域名、端口
	PrettyUrl     string            `json:"pretty_url"` //不包括协议、域名和端口
	RawUrl        string            `json:"raw_url"`
	Scheme        string            `json:"scheme"`
	HttpVersion   string            `json:"http_version"`
	Method        string            `json:"method"`
	Port          int               `json:"port"`
	RequestHost   string            `json:"request_host"`
	Project       string            `json:"project"`
	ResponseCode  int               `json:"response_code"`
	Header        map[string]string `json:"header"`
	SplitUrl      string            `json:"split_url"`
	StandardMd5   string            `json:"standard_md5"`
	UrlParameter  string            `json:"url_parameter,omitempty"`
	BodyParameter string            `json:"body_parameter,omitempty"`
}
