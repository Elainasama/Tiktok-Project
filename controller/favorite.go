package controller

import (
	"TikTokLite/common"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FavoriteAction 用来添加和删除点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	userId, err := common.VerifyToken(token)
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}

	response, err := service.FavoriteActionService(userId, videoId, int32(actionType))
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}

// FavoriteList 获取点赞过的所有视频
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	_, err := common.VerifyToken(token)
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}
	uid, _ := strconv.ParseInt(userId, 10, 64)

	response, err := service.FavoriteListService(uid)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}
