package controllers

import (
	"LiteService/app/engine"
	"net/http"

	"github.com/gin-gonic/gin"
)

type res struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
}
type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (u *HelloController) Say() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, res{
			Data: nil,
			Msg:  "ok",
			Code: 0,
		})
	}
}
func (u *HelloController) Build(g *engine.Engine) {
	r := g.Group("/hello")
	r.POST("/", u.Say())
}
