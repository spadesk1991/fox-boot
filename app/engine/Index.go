package engine

import (
	"LiteService/app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IService interface {
	Build(engine *Engine)
}

type IRegistry interface {
	Reg()
}

type IUnRegistry interface {
	UnReg()
}

type Engine struct {
	*gin.Engine
}

func NewEngine() *Engine {
	e := gin.Default()
	e.NoMethod(middleware.HandleNotFound)
	e.NoRoute(middleware.HandleNotFound)
	e.Use(middleware.ErrHandler())
	return &Engine{Engine: e}
}

func (service *Engine) Mount(controllers ...IService) *Engine {
	for _, controller := range controllers {
		controller.Build(service)
	}

	return service
}

type returnResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func (r *returnResponse) json(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func (service *Engine) handle(httpMethod, relativePath string, handlers ...interface{}) {
	arr := make([]gin.HandlerFunc, 0)
	for _, handler := range handlers {
		switch handler.(type) {
		case func(*gin.Context):
			arr = append(arr, handler.(func(*gin.Context)))
		case func() (string, error):
			f := func(c *gin.Context) {
				res, err := handler.(func() (string, error))()
				if err != nil {
					panic(err)
				}
				c.JSON(http.StatusOK, res)
				r := &returnResponse{
					Code:   0,
					Msg:    "ok",
					Result: res,
				}
				r.json(c)
			}
			arr = append(arr, f)
		case func() (int, error):
			f := func(c *gin.Context) {
				res, err := handler.(func() (int, error))()
				if err != nil {
					panic(err)
				}
				r := returnResponse{
					Code:   0,
					Msg:    "ok",
					Result: res,
				}
				r.json(c)
			}
			arr = append(arr, f)
		case func() (interface{}, error):
			f := func(c *gin.Context) {
				res, err := handler.(func() (interface{}, error))()
				if err != nil {
					panic(err)
				}
				r := returnResponse{
					Code:   0,
					Msg:    "ok",
					Result: res,
				}
				r.json(c)
			}
			arr = append(arr, f)
		case func() (bool, error):
			f := func(c *gin.Context) {
				res, err := handler.(func() (bool, error))()
				if err != nil {
					panic(err)
				}
				r := returnResponse{
					Code:   0,
					Msg:    "ok",
					Result: res,
				}
				r.json(c)
			}
			arr = append(arr, f)
		default:
			panic("不支持的controller函数类型")
		}
	}
	service.Engine.Handle(httpMethod, relativePath, arr...)
}

func (service *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *Engine {
	service.Engine.RouterGroup = *service.Engine.Group(relativePath, handlers...)
	return service
}

func (service *Engine) Use(middleware ...gin.HandlerFunc) *Engine {
	service.Engine.Use(middleware...)
	return service
}

func (service *Engine) POST(relativePath string, handlers ...interface{}) {
	service.handle(http.MethodPost, relativePath, handlers...)
}

func (service *Engine) GET(relativePath string, handlers ...interface{}) {
	service.handle(http.MethodGet, relativePath, handlers...)
}

func (service *Engine) PUT(relativePath string, handlers ...interface{}) {
	service.handle(http.MethodPut, relativePath, handlers...)
}

func (service *Engine) DELETE(relativePath string, handlers ...interface{}) {
	service.handle(http.MethodDelete, relativePath, handlers...)
}

func (service *Engine) Registry(reg IRegistry) *Engine {
	reg.Reg()
	return service
}

func (service *Engine) UnRegistry(unReg IUnRegistry) *Engine {
	unReg.UnReg()
	return service
}

func (service *Engine) Run(addr ...string) {
	err := service.Engine.Run(addr...)
	if err != nil {
		panic(err)
	}
}
