package engine

import (
	"LiteService/app/common"
	"LiteService/app/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type IService interface {
	Build(engine *Engine)
}
type Engine struct {
	*gin.Engine
	prefix string
}

func NewEngine(prefix string) *Engine {
	e := gin.Default()
	e.NoMethod(middleware.HandleNotFound)
	e.NoRoute(middleware.HandleNotFound)
	e.Use(middleware.ErrHandler())
	e.GET("/check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r := e.Group(prefix)
	e.RouterGroup = *r
	return &Engine{Engine: e, prefix: prefix}
}

func (service *Engine) Mount(controllers ...IService) *Engine {
	for _, controller := range controllers {
		controller.Build(service)
	}
	if os.Getenv("GIN_MODE") != "" {
		service.registry()
	}
	return service
}

func (service *Engine) registry() {
	key := fmt.Sprintf("http://172.17.0.1:%s", os.Getenv("port"))
	path := fmt.Sprintf("%s/*", service.prefix)
	apis := map[string]interface{}{key: []string{path}}
	res, err := common.Post("http://172.17.0.1:88/sign", nil, apis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("registry router ok, at %v\n", string(res))
}
