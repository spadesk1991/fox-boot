package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Cfg *Config

type Config struct {
	Port          int    `yaml:"port"`
	Debug         bool   `yaml:"debug"`
	MysqlUri      string `yaml:"mysqlUri"`
	MongodbUri    string `yaml:"mongodbUri"`
	RedisAddr     string `yaml:"redisAddr"`
	RedisPassword string `yaml:"redisPassword"`
	RedisDb       int    `yaml:"redisDb"`
	RedisTimeout  int    `yaml:"redisTimeout"`
}

func init() {
	goEnv := os.Getenv("GIN_MODE")
	if goEnv == "" {
		goEnv = "default"
	}
	//dir, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("config/%s.yaml", goEnv))
	if err != nil {
		panic(err)
	}
	Cfg = new(Config)
	err = yaml.Unmarshal(yamlFile, Cfg)
	if err != nil {
		panic(err)
	}
}
