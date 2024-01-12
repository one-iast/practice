package http

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"iast-demo/entity"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func SendRequest(request *http.Request, proxyUrl *url.URL) (*http.Response, error) {
	//proxyAddr := "http://127.0.0.1:8080"
	//proxyUrl, err := url.Parse(proxyAddr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxyUrl),
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	client := http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GenHttpRequest(c *entity.HttpData) *http.Request {
	if c.Api != "" {
		var (
			body io.Reader
			api  string
		)
		bodyData := c.BodyData
		if bodyData != nil {
			bodyType := reflect.TypeOf(bodyData)
			if bodyType.String() == "string" {
				strBody := bodyData.(string)
				body = strings.NewReader(strBody)
			} else {
				marshal, err := json.Marshal(bodyData)
				if err != nil {
					log.Fatal(err)
				}
				body = strings.NewReader(string(marshal))
			}
		}
		if c.UrlData != "" {
			//?拼接的参数
			if strings.Contains(c.UrlData, "=") {
				api = c.Api + "?" + c.UrlData
			} else {
				//restful 参数
				api = c.Api + "/" + c.UrlData
			}
		} else {
			api = c.Api
		}
		req, _ := http.NewRequest(c.HttpMethod, api, body)
		if len(c.Header) != 0 {
			for k, v := range c.Header {
				req.Header.Add(k, v)
			}
		}
		return req
	}
	return nil
}

func GetResponseBody(resp *http.Response) []byte {
	body, _ := io.ReadAll(resp.Body)
	err := resp.Body.Close()
	if err != nil {
		return nil
	}
	return body
}

func DoRequest(httpData *entity.HttpData, proxyUrls ...*url.URL) ([]byte, int, error) {
	var proxyUrl *url.URL
	if len(proxyUrls) != 0 {
		proxyUrl = proxyUrls[0]
	}

	request := GenHttpRequest(httpData)
	resp, err := SendRequest(request, proxyUrl)
	if resp == nil {
		if err == nil {
			err = errors.New("nil resp")
		}
		return nil, 0, err
	}
	body := GetResponseBody(resp)
	return body, resp.StatusCode, err
}
