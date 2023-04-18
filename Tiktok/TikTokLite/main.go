package main

import (
	"TikTokLite/common"
	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	defer common.CloseDataBase()
	//defer common.CloseRabbitMq()
	//go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Init() {
	common.InitConfig()
	common.DBInit()
	common.InitMinIO()
	//common.InitRabbitMq()
}
