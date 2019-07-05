package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shorturl/model"
	"strconv"
)

var User = &UserController{}

type UserController struct {
	BaseController
}

func (i *UserController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println("username: %s, password: %s", username, password)

	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"username": username,
		"password": password,
	})
}

func (i *UserController) UserInfo(c *gin.Context) {
	page := c.DefaultQuery("page", "0")

	var pageInt, _ = strconv.Atoi(page)
	var limit = 10
	var offset = pageInt * limit

	//var user model.User

	users := make([]*model.User, 0)

	model.DB.Table("user").Where("weight >= ?", 0).Offset(offset).Limit(limit).Order("id asc").Find(&users)

	i.success(c, gin.H{
		"user": users,
	})
}
