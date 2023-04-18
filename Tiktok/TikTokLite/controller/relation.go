package controller

import (
	"TikTokLite/common"
	. "TikTokLite/message"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	userId, err := common.VerifyToken(token)

	if userId == toUserId {
		c.JSON(http.StatusOK, DouyinRelationActionResponse{Response: Response{StatusCode: 1, StatusMsg: "can not follow yourself"}})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusOK, DouyinRelationActionResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	response, err := service.RelationActionService(userId, toUserId, int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, DouyinRelationActionResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, *response)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response, err := service.GetRelationFollowListService(userId)
	if err != nil {
		c.JSON(http.StatusOK, DouyinRelationFollowListResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, *response)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response, err := service.GetRelationFollowerListService(userId)
	if err != nil {
		c.JSON(http.StatusOK, DouyinRelationFollowerListResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, *response)
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response, err := service.GetRelationFriendService(userId)
	if err != nil {
		c.JSON(http.StatusOK, DouyinRelationFollowerListResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, *response)
}
