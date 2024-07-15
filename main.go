package main

import (
	"log"
	"wxpusher/api"
	"wxpusher/config"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("读取环境配置出错:", err)
	}
	
	router := gin.Default()
	api.NewWxRouter(router)
	log.Fatal(router.Run(":8080"))
}
