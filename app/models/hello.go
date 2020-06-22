package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;type:varchar(20)"`
	age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

func (u *Users) TableName() string {
	return "user1"
}

func (u *Users) BeforeCreate() (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *Users) BeforeUpdate() (err error) {
	u.UpdatedAt = time.Now()
	return
}
