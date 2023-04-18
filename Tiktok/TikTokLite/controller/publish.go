package controller

import (
	"TikTokLite/common"
	. "TikTokLite/message"
	"TikTokLite/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	conf := common.GetConfig()

	token := c.PostForm("token")
	title := c.PostForm("title")
	userid, err := common.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	saveFile := filepath.Join(conf.Path.VideoPath, finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	log.Println("Upload File: ", saveFile)

	response, err := service.PublishVideo(userid, saveFile, title)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, *response)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	uid := c.Query("user_id")
	id, _ := strconv.ParseInt(uid, 10, 64)
	userId, err := common.VerifyToken(c.Query("token"))
	if err != nil || userId != id {
		c.JSON(200, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	PublishListResponse, err := service.GetPublishList(id)
	if err != nil {
		c.JSON(200, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(200, *PublishListResponse)
}
