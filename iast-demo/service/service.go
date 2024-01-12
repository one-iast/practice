package service

import (
	"github.com/redis/go-redis/v9"
	"iast-demo/entity/config"
	"log"
)

// NewRedis 创建redis客户端
func NewRedis() (*redis.Client, func(), error) {
	dsn := config.CFG.Cache.Redis.DSN()
	// 有两种创建client的方式
	// See: https://redis.uptrace.dev/guide/go-redis.html#connecting-to-redis-server
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(opt)
	return rdb, func() {
		err := rdb.Close()
		if err != nil {
			return
		}
	}, err
}
