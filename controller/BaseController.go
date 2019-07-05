package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/model"
)

type BaseController struct {}

func (*BaseController) success(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": model.Success,
		"msg":  "ok",
		"data": data,
	})
}

func (*BaseController) failed(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
	c.Abort()
}