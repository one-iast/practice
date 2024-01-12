package main

import (
	"iast-demo/common"
	"iast-demo/entity/config"
	"iast-demo/service/burpsuite"
	"iast-demo/service/sqlmap"
	"iast-demo/util"
)

func main() {
	common.ShowVersion()

	config.LoadConfig(".")

	if config.CFG.Sqlmap.Enable {
		go util.Run(new(sqlmap.Runner))
	}
	if config.CFG.Burpsuite.Enable {
		go util.Run(new(burpsuite.Runner))
	}
	util.KeepAlive()
}
