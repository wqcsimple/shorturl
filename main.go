package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"shorturl/model"
	"shorturl/routers"
	"time"
)

func init() {
	dir, _ := os.Getwd()
	file := dir + "/conf/app.conf"

	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Panicf("conf file [%s]  not found!", file)
	}
	model.Conf.InitConfig(file)
}

func main() {
	gin.SetMode(gin.DebugMode)
	if model.Conf.AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	//logFile, _ := os.Create("storage/logs/access.log")
	//gin.DefaultWriter = logFile

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		//定制日志格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())


	err := logs.SetLogger(logs.AdapterFile, `{"filename":"storage/logs/app.log"}`)
	if err != nil {
		log.Fatalln("Log init failed, error: " + err.Error())
	}

	routers.Route(router)

	err = router.Run(":" + model.Conf.AppPort)
	if err != nil {
		log.Fatalln("Server start failed, error: " + err.Error())
	}
}
