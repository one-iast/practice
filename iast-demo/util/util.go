package util

import (
	"context"
	"encoding/json"
	"fmt"
	"iast-demo/common"
	"iast-demo/entity"
	"iast-demo/entity/config"
	"iast-demo/service"
	"log"
	"net/url"
	"strings"
)

// flow -> packet
func PacketFlow() {
	rdb, f, err := service.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer f()
	ctx := context.Background()
	iter := rdb.Scan(ctx, 0, "flow*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		resultList, err := rdb.RPopCount(ctx, key, config.CFG.Stage.PacketSize).Result()
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range resultList {
			flow := &entity.Flow{}
			json.Unmarshal([]byte(v), flow)
			log.Printf("packet flow【%s %s】", key, flow.RawUrl)

			urlWithPrefix := fmt.Sprintf("%s://%s", flow.Scheme, flow.RequestHost)
			if flow.Port != 80 && flow.Port != 443 {
				urlWithPrefix = fmt.Sprintf("%s:%d", urlWithPrefix, flow.Port)
			}
			flow.Url = urlWithPrefix
			prettyUrl := strings.ReplaceAll(flow.RawUrl, flow.Url, "")
			flow.PrettyUrl = prettyUrl

			//header的key转小写，方便获取对应值
			lowerHeaderMap := common.LowerMapKey(flow.Header)

			//sqlmap数据封装
			sqlmapFlow := PackageSqlmapFlow(flow, lowerHeaderMap)
			rdb.LPush(ctx, "packet:sqlmap:"+flow.RequestHost, sqlmapFlow)

			//burp数据封装，直接复制一份原来的flow
			marshal, _ := json.Marshal(flow)
			//分开不同的用户
			rdb.LPush(ctx, "packet:burpsuite:"+flow.RequestHost+":"+flow.Username, marshal)
		}
	}
}

func PackageSqlmapFlow(flow *entity.Flow, lowerHeaderMap map[string]string) []byte {
	//封装sqlmap数据
	sqlmapData := &entity.SqlmapTaskFlow{
		Url:    flow.RawUrl,
		Method: flow.Method,
		Data:   flow.BodyParameter,
		//请求头中有Content-Length时，sqlmap扫描会失败，需要去除
		Headers: common.Map2StringLine(flow.Header, "content-length"),
		Cookie:  lowerHeaderMap["cookie"],
		Agent:   lowerHeaderMap["user-agent"],
		Host:    lowerHeaderMap["host"],
		Referer: lowerHeaderMap["referer"],
	}
	sqlmapPacketData := &entity.SqlmapPacketData{
		SqlmapTaskFlow: sqlmapData,
		Flow:           flow,
	}
	marshal, _ := json.Marshal(sqlmapPacketData)
	return marshal
}

func KeepAlive() {
	var ch chan int
	<-ch
}

func GetHttpProxy(proxyInfos ...string) *url.URL {
	proxyAddr := "http://127.0.0.1:8080"
	if len(proxyInfos) == 3 {
		proxyAddr = fmt.Sprintf("%s://%s:%s", proxyInfos[0], proxyInfos[1], proxyInfos[2])
	}
	proxyUrl, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatalf("parse proxy address failed: %s", err)
	}
	return proxyUrl
}

func Run(scanner entity.Scanner) {
	scanner.Run()
}
