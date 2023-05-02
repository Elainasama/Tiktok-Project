package main

import (
	"TikTokLite/common"
	"TikTokLite/logger"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// go tool pprof http://localhost:8080/debug/pprof/goroutine?second=120

func main() {
	Init()
	defer Close()

	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	pprof.Register(r)

	initRouter(r)

	r.Run()
}

func Init() {
	common.InitConfig()
	common.DBInit()
	common.InitFilter()
	common.InitMinIO()
	common.InitRedis()
	logger.InitLogger()
	//common.InitRabbitMq()
}

func Close() {
	common.CloseDataBase()
	//common.CloseRabbitMq()
	common.CloseRedis()
	logger.Sync()
}
