package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

// DB连接
var DB *gorm.DB

// redis连接
var Redis *redis.Pool

func BaseInit() {
	initDB()
	initRedis()
}

func initDB() {
	var db *gorm.DB
	url := Conf.DBUser + ":" + Conf.DBPass + "@tcp(" + Conf.DBHost + ":" + Conf.DBPort + ")/" + Conf.DBDatabase +
		"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(Conf.DBDialect, url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	DB = db
}

func initRedis()  {
	server := Conf.RedisHost + ":" + Conf.RedisPort

	redisPool := &redis.Pool{
		MaxIdle:     Conf.RedisMaxIdle,
		MaxActive:   Conf.RedisMaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialPassword(Conf.RedisPass))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	Redis = redisPool
}