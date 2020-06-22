package models

import "LiteService/app/dao"

func init() {
	dao.GetMysqlDB().AutoMigrate(Users{})
}
