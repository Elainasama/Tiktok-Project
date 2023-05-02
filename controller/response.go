package controller

import (
	"TikTokLite/logger"
	. "TikTokLite/message"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var tokenError = "User doesn't exist"
var failCode int32 = 1

// 重构点 把频繁使用的发送Json代码封装成函数直接调用 减少代码冗余
// 抽象成两句代码
// ResponseFail(c, err.Error())
// ResponseSuccess(c, *response)

func ResponseSuccess(c *gin.Context, v interface{}) {
	logger.Infof("ResponseSuccess: %v", reflect.TypeOf(v))
	c.JSON(http.StatusOK, v)
}

func ResponseFail(c *gin.Context, msg string) {
	logger.Errorf("ResponseFail:", msg)
	c.JSON(http.StatusOK, Response{
		StatusCode: failCode,
		StatusMsg:  msg,
	})
}
