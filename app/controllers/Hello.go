package controllers

import (
	"LiteService/app/engine"
)

type res struct {
	Data interface{} `json:"data" gorm:"column:data"`
	Msg  string      `json:"msg" gorm:"column:msg"`
	Code int         `json:"code" gorm:"column:code"`
}
type helloController struct{}

func NewHelloController() *helloController {
	return &helloController{}
}

func (u *helloController) Say() (string, error) {
	return "hello", nil
}

func (u *helloController) Build(g *engine.Engine) {
	r := g.Group("/hello")
	r.GET("/", u.Say)
}
