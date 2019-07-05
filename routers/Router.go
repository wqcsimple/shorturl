package routers

import (
	"github.com/gin-gonic/gin"
	"shorturl/controller"
)

func Route(Router *gin.Engine) {

	//Router.GET("/:code", Index.Path)

	user := Router.Group("client/1/user")
	{
		user.POST("login", controller.Login)
		user.GET("info", controller.UserInfo)
	}

	file := Router.Group("file")
	{
		file.POST("upload", controller.Upload)
	}


	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})

	//router.GET("redirect", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "http://whis.wang")
	//})
}
