package main

import (
	"LiteService/app/controllers"
	"LiteService/app/engine"
	"LiteService/config"
	"fmt"
)

func main() {
	router := engine.NewEngine("/api/demo")
	router.Group("/api")
	// 加载路由
	router.Mount(
		controllers.NewHelloController(),
	)

	err := router.Run(fmt.Sprintf(":%d", config.Cfg.Port))
	if err != nil {
		panic(err)
	}
}
