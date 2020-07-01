package dao

import (
	"LiteService/config"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", config.GetCfg().MysqlUri)
	if err != nil {
		panic(err)
	}
	if config.GetCfg().Gin_mode != "release" { // 生产环境关闭log
		db.LogMode(true)
	}
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}

func GetMysqlDB() *gorm.DB {
	return db
}
