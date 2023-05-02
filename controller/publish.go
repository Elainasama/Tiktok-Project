package controller

import (
	"TikTokLite/common"
	"TikTokLite/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"strconv"
)

// Publish 支持用户上传视频
func Publish(c *gin.Context) {
	conf := common.GetConfig()

	token := c.PostForm("token")
	title := c.PostForm("title")
	userid, err := common.VerifyToken(token)
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	saveFile := filepath.Join(conf.Path.VideoPath, finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		ResponseFail(c, err.Error())
		return
	}

	log.Println("Upload File: ", saveFile)

	response, err := service.PublishVideo(userid, saveFile, title)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *response)
}

// PublishList 支持用户获取其他用户已上传的视频列表
func PublishList(c *gin.Context) {
	uid := c.Query("user_id")
	id, _ := strconv.ParseInt(uid, 10, 64)
	_, err := common.VerifyToken(c.Query("token"))
	if err != nil {
		ResponseFail(c, tokenError)
		return
	}
	PublishListResponse, err := service.GetPublishList(id)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}

	ResponseSuccess(c, *PublishListResponse)
}
