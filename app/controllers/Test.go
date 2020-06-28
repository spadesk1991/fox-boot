package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (u *TestController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		//testService := new(service.TestService)
		c.JSON(http.StatusOK, gin.H{"err_code": 0, "data": "ok"})
	}
}
