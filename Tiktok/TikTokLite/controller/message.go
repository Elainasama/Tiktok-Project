package controller

import (
	"TikTokLite/common"
	. "TikTokLite/message"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	userid, err := common.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	toUserid, _ := strconv.ParseInt(toUserId, 10, 64)

	response, err := service.MessageSendService(userid, toUserid, content)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, *response)
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserIdString := c.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64)
	PreMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	userid, err := common.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	response, err := service.GetMessageList(userid, toUserId, PreMsgTime)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, *response)
}
