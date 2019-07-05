package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/common"
)

var File = &FileController{}

type FileController struct {
	BaseController
}

func (i * FileController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	common.CheckError(err)

	fmt.Println(file.Filename)

	//c.SaveUploadedFile(file, "./test.sql")

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "filename: " + file.Filename,
	})
}
