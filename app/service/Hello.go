package service

import (
	"LiteService/app/dao"
	"LiteService/app/models"

	"github.com/jinzhu/gorm"
)

type HelloService struct{}

func (s HelloService) CreateHello() (err error) {
	return dao.GetMysqlDB().Transaction(func(tx *gorm.DB) (e error) {
		user := models.Users{
			Name: "张三",
			Sex:  13,
		}
		if e = tx.Create(&user).Error; e != nil {
			return
		}
		if e = tx.Unscoped().Delete(&models.Users{}, "sex =13").Error; e != nil {
			return
		}
		return
	})
}
