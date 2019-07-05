package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shorturl/model"
)



func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println("username: %s, password: %s", username, password)

	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"username": username,
		"password": password,
	})
}

func UserInfo(c *gin.Context) {
	userId := c.Query("id")

	log.Println("id -> ", userId)

	var user model.User
	model.DB.First(&user)

	var rst = make(map[string]model.User)
	rst["user"] = user

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rst,
	})
}


