package main

import (
	"log"
	"wxpusher/api"
	"wxpusher/config"
	"wxpusher/pkg/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("读取环境配置出错:", err)
	}
	rdb := redis.InitDB()
	router := gin.Default()
	router.Static("/public", "./public")
	api.NewWxRouter(router, rdb)
	log.Fatal(router.Run(":8080"))
}
