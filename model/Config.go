package model

import (
	"github.com/astaxie/beego/config"
)

type App struct {
	AppName string
	AppPort string
	AppEnv  string
	AppUrl  string

	DBDialect      string
	DBHost         string
	DBPort         string
	DBDatabase     string
	DBUser         string
	DBPass         string

	RedisHost      string
	RedisPort      string
	RedisPass      string
	RedisMaxIdle   int
	RedisMaxActive int
}

var Conf *App

func init() {
	Conf = &App{}
}

func (App) InitConfig(file string)  {
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		panic(err)
	}

	Conf.AppName = conf.DefaultString("APP_NAME", "")
	Conf.AppPort = conf.DefaultString("APP_PORT", "8080")
	Conf.AppEnv = conf.DefaultString("APP_ENV", "dev")
	Conf.AppUrl = conf.DefaultString("APP_URL", "http://127.0.0.1:8080/")

	Conf.DBDialect = conf.DefaultString("DB_Dialect", "mysql")
	Conf.DBHost = conf.DefaultString("DB_HOST", "db")
	Conf.DBPort = conf.DefaultString("DB_PORT", "3306")
	Conf.DBDatabase = conf.DefaultString("DB_DATABASE", "short")
	Conf.DBUser = conf.DefaultString("DB_USERNAME", "root")
	Conf.DBPass = conf.DefaultString("DB_PASSWORD", "7")

	Conf.RedisHost = conf.DefaultString("REDIS_HOST", "redis")
	Conf.RedisPort = conf.DefaultString("REDIS_PORT", "6379")
	Conf.RedisPass = conf.DefaultString("REDIS_PASS", "")
	Conf.RedisMaxIdle = conf.DefaultInt("REDIS_MAX_IDLE", 8)
	Conf.RedisMaxActive = conf.DefaultInt("REDIS_MAX_ACTIVE", 100)

	BaseInit()
}