package controller

import (
	"TikTokLite/common"
	. "TikTokLite/message"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentAction no practical effect, just check if token is valid
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
		c.JSON(http.StatusOK, DouyinCommentActionResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	response, err := service.CommentActionService(userid, videoId, int32(actionType), text, commentId)
	if err != nil {
		c.JSON(http.StatusOK, DouyinCommentActionResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	_, err := common.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusOK, DouyinCommentActionResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	response, err := service.GetCommentListService(videoId)
	if err != nil {
		c.JSON(http.StatusOK, DouyinCommentActionResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, *response)
}
