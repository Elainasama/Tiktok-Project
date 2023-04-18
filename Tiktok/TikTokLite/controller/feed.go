package controller

import (
	. "TikTokLite/message"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed 视频流
func Feed(c *gin.Context) {
	currentTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		currentTime = time.Now().Unix()
	}
	//userid, _ := c.Get("UserId")
	//uid := userid.(int64)
	videoListResponse, err := service.GetVideoListResponse(currentTime)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, *videoListResponse)
}
