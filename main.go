package main

import (
	"LiteService/app/engine"
	"LiteService/app/service"
	"LiteService/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := engine.NewEngine()
	r.GET("/check", func(c *gin.Context) (string, error) {
		return "ok", nil
	})
	r.Group(config.GetCfg().Prefix).
		Mount().
		Register(service.Reg).
		Run()
}
