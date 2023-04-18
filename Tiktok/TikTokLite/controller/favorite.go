package controller

import (
	"TikTokLite/common"
	. "TikTokLite/message"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	userId, err := common.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusOK, DouyinFavoriteActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}
	response, err := service.FavoriteActionService(userId, videoId, int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, DouyinFavoriteActionResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, *response)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	Id, err := common.VerifyToken(token)
	uid, _ := strconv.ParseInt(userId, 10, 64)

	log.Println("FavoriteList", Id, uid)

	if err != nil || Id != uid {
		c.JSON(http.StatusOK, DouyinFavoriteListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "token error"},
			VideoList: nil,
		})
		return
	}
	response, err := service.FavoriteListService(uid)
	if err != nil {
		c.JSON(http.StatusOK, DouyinFavoriteListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, *response)
}
