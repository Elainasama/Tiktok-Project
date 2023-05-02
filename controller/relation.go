package controller

import (
	"TikTokLite/common"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// RelationAction 支持用户完成关注与取关行为
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	userId, err := common.VerifyToken(token)

	if userId == toUserId {
		ResponseFail(c, "can not follow yourself")
		return
	}

	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	response, err := service.RelationActionService(userId, toUserId, int32(actionType))
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}

// FollowList 支持用户获取关注列表
func FollowList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response, err := service.GetRelationFollowListService(userId)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}

// FollowerList 支持用户获取粉丝列表
func FollowerList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response, err := service.GetRelationFollowerListService(userId)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}

// FriendList 支持用户获取互相关注列表
func FriendList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response, err := service.GetRelationFriendService(userId)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}
