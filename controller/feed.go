package controller

import (
	"TikTokLite/common"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed 获取视频流 用于首页展示 Id和currentTime是可选参数 保留默认参数
func Feed(c *gin.Context) {
	currentTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		currentTime = time.Now().Unix()
	}
	var userId int64
	token := c.Query("token")
	if token == "" {
		userId = -1
	} else {
		userId, err = common.VerifyToken(token)
		if err != nil {
			ResponseFail(c, tokenError)
			return
		}
	}

	videoListResponse, err := service.GetVideoListResponse(userId, currentTime)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, *videoListResponse)
}
