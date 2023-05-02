package controller

import (
	"TikTokLite/common"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// MessageAction 用户发送消息
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	userid, err := common.VerifyToken(token)
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}
	toUserid, _ := strconv.ParseInt(toUserId, 10, 64)

	response, err := service.MessageSendService(userid, toUserid, content)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}

// MessageChat 用户获取聊天记录
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserIdString := c.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64)
	PreMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	userid, err := common.VerifyToken(token)
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}
	response, err := service.GetMessageList(userid, toUserId, PreMsgTime)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}
