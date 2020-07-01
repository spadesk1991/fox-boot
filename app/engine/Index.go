package engine

import (
	"LiteService/app/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

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
	unRegistryFunc RegistryFunc
}

type RegistryFunc func()

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
		case func(c *gin.Context):
			arr = append(arr, handler.(func(c *gin.Context)))
		case func(c *gin.Context) (string, error):
			f := func(c *gin.Context) {
				res, err := handler.(func(c *gin.Context) (string, error))(c)
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
		case func(c *gin.Context) (int, error):
			f := func(c *gin.Context) {
				res, err := handler.(func(c *gin.Context) (int, error))(c)
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
		case func(c *gin.Context) (interface{}, error):
			f := func(c *gin.Context) {
				res, err := handler.(func(c *gin.Context) (interface{}, error))(c)
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
		case func(c *gin.Context) (bool, error):
			f := func(c *gin.Context) {
				res, err := handler.(func(c *gin.Context) (bool, error))(c)
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

func (service *Engine) Register(reg RegistryFunc) *Engine {
	// 服务注册
	reg()
	return service
}

func (service *Engine) Deregister(f RegistryFunc) *Engine {
	service.unRegistryFunc = f
	return service
}

func (service *Engine) Run(addr ...string) {
	go func(service *Engine) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			s := <-c
			logrus.Infof("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				logrus.Info("application exit")
				// 服务反注册
				if service.unRegistryFunc != nil {
					service.unRegistryFunc()
				}
				time.Sleep(time.Second)
				os.Exit(0)
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}(service)
	err := service.Engine.Run(addr...)
	if err != nil {
		panic(err)
	}
}
