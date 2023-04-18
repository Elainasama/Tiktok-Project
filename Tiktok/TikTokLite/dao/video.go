package dao

import (
	"TikTokLite/common"
	. "TikTokLite/model"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func GetVideoList(currentTime int64) ([]Video, error) {
	db := common.GetDB()
	var res []Video
	err := db.Where("publish_time < ?", currentTime).Order("publish_time DESC").Limit(30).Find(&res).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return res, err
	}
	// 后续可以加个缓存
	for i, video := range res {
		author, err := GetUserInfo(video.AuthorId)
		if err != nil {
			return res, err
		}
		res[i].Author = author
	}
	return res, nil
}

func GetPublishList(userid int64) ([]Video, error) {
	db := common.GetDB()
	author, err := GetUserInfo(userid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	var res []Video
	err = db.Where("author_id = ?", userid).Find(&res).Error
	for i := range res {
		res[i].Author = author
	}
	return res, err
}

func InsertVideo(userid int64, title string, videoUrl string, imageUrl string) error {
	db := common.GetDB()
	video := Video{
		AuthorId:      userid,
		PlayUrl:       videoUrl,
		CoverUrl:      imageUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		PublishTime:   time.Now().Unix(),
		Title:         title,
	}
	err := db.Create(&video).Error
	if err != nil {
		return errors.New("Insert Video Failed")
	}
	return nil
}

func ChangeVideoCount(videoId int64, countType string, diff int64) {
	db := common.GetDB()
	var video Video
	err := db.Where("video_id = ?", videoId).First(&video).Error
	if err != nil {
		log.Fatalln(err)
	}

	switch countType {
	case "favorite_count":
		video.FavoriteCount += diff
	case "comment_count":
		video.CommentCount += diff
	}

	err = db.Save(&video).Error
	if err != nil {
		log.Fatalln(err)
	}
}

func GetVideoAuthorId(videoId int64) int64 {
	// 可以增加缓存
	db := common.GetDB()
	var video Video
	err := db.Where("video_id = ?", videoId).Select("author_id").First(&video).Error
	if err != nil {
		log.Fatalln(err)
	}
	return video.AuthorId
}

func VideoIsExist(videoId int64) bool {
	db := common.GetDB()
	var video Video
	err := db.Where("video_id = ?", videoId).First(&video).Error
	return err == nil
}
