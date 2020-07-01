package config

import (
	"os"
	"strconv"
)

var cfg *Config

type Config struct {
	Prefix        string
	Port          int
	Gin_mode      string
	MysqlUri      string
	MongodbUri    string
	RedisAddr     string
	RedisPassword string
	RedisDb       int
}

func init() {
	if err := newConfig(); err != nil {
		panic(err)
	}
}

func newConfig() (err error) {
	cfg = new(Config)
	if cfg.Port, err = strconv.Atoi(os.Getenv("port")); err != nil {
		return
	}
	cfg.Prefix = os.Getenv("prefix")
	cfg.MysqlUri = os.Getenv("mysqlUri")
	cfg.MongodbUri = os.Getenv("mongodbUri")
	cfg.Gin_mode = os.Getenv("GIN_MODE")
	cfg.RedisAddr = os.Getenv("redisAddr")
	cfg.RedisPassword = os.Getenv("redisPassword")
	if cfg.RedisDb, err = strconv.Atoi(os.Getenv("redisDb")); err != nil {
		return
	}
	return
}

func GetCfg() *Config {
	return cfg
}
