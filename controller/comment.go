package controller

import (
	"TikTokLite/common"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CommentAction 用来添加和删除评论操作
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	text := c.Query("comment_text")
	var commentId int64
	commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	if err != nil {
		commentId = 0
	}
	userid, err := common.VerifyToken(token)
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}
	response, err := service.CommentActionService(userid, videoId, int32(actionType), text, commentId)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}

// CommentList 获取视频下的所有评论
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	_, err := common.VerifyToken(token)
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}
	response, err := service.GetCommentListService(videoId)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}
