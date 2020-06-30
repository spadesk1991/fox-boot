package dao

import (
	"LiteService/config"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisAddr,
		Password: config.Cfg.RedisPassword, // no password set
		DB:       config.Cfg.RedisDb,       // use default DB
	})
	fmt.Println(rdb.Ping(context.TODO()))
}

func GetRedisDB() *redis.Client {
	return rdb
}
