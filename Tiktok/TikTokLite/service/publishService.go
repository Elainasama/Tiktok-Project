package service

import (
	"TikTokLite/common"
	"TikTokLite/dao"
	. "TikTokLite/message"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func GetPublishList(userid int64) (*VideoListResponse, error) {
	videoList, err := dao.GetPublishList(userid)
	if err != nil {
		return nil, err
	}
	return &VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: messageVideoList(videoList),
	}, nil
}

func PublishVideo(userid int64, saveFile string, title string) (*DouyinPublishActionResponse, error) {
	m := common.GetMinIO()
	videoUrl, err := m.UploadFile(saveFile, strconv.FormatInt(userid, 10), "video")
	if err != nil {
		return nil, err
	}

	imgPath, err := GetImageFile(saveFile)
	if err != nil {
		return nil, err
	}

	imageUrl, err := m.UploadFile(imgPath, strconv.FormatInt(userid, 10), "image")
	if err != nil {
		return nil, err
	}

	err = dao.InsertVideo(userid, title, videoUrl, imageUrl)
	if err != nil {
		return nil, err
	}

	return &DouyinPublishActionResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "Video Upload Success",
		}}, nil
}

func GetImageFile(videoPath string) (string, error) {
	conf := common.GetConfig()
	sp := strings.Split(videoPath, "\\")
	videoName := sp[len(sp)-1]
	imageName := string(videoName[:len(videoName)-3]) + "jpg"
	imagePath := filepath.Join(conf.Path.ImagePath, imageName)
	log.Println("Image Path: ", imagePath, "Video Path: ", videoPath, "Image Name: ", imageName)
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vframes", "1", "-an", "-s", "640x360", "-ss", "00:00:01", imagePath)
	err := cmd.Run()
	if err != nil {
		log.Println("Image Generate Fail", err)
		return "", err
	}
	return imagePath, nil
}
