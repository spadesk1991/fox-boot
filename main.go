package main

import (
	"LiteService/app/controllers"
	"LiteService/app/engine"
	"LiteService/app/service"
	"LiteService/config"
	"fmt"
)

func main() {
	prefix := "/api/demo"
	r := engine.NewEngine()
	r.GET("/check", func() (string, error) {
		return "ok", nil
	})
	r.Group(prefix).
		Mount(
			controllers.NewHelloController(),
		).
		Registry(service.NewRegService(prefix)).
		Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
