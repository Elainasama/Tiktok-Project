package dao

import (
	"TikTokLite/common"
	. "TikTokLite/model"
	"errors"
	"github.com/jinzhu/gorm"
)

func InsertFavoriteAction(userId int64, videoId int64) error {
	db := common.GetDB()

	if VideoIsExist(videoId) == false {
		return errors.New("video did not Exist")
	}

	err := db.Where("user_id = ? and video_id = ?", userId, videoId).First(&Favorite{}).Error
	if err == gorm.ErrRecordNotFound {
		favorite := Favorite{
			UserId:  userId,
			VideoId: videoId,
		}
		err := db.Create(&favorite).Error
		if err != nil {
			return errors.New("insert Favorite Failed")
		}
		authorId := GetVideoAuthorId(videoId)
		ChangeVideoCount(videoId, "favorite_count", 1)
		go ChangeUserCount(userId, "favorite_count", 1)
		go ChangeUserCount(authorId, "favorited_count", 1)
		return nil
	}
	return errors.New("favorite Record Exist")
}

func DeleteFavoriteAction(userId int64, videoId int64) error {
	db := common.GetDB()

	if VideoIsExist(videoId) == false {
		return errors.New("video did not Exist")
	}

	var favorite Favorite
	err := db.Where("user_id = ? and video_id = ?", userId, videoId).First(&favorite).Error
	if err != nil {
		return errors.New("record did not Exist")
	}
	err = db.Delete(&favorite).Error
	if err != nil {
		return errors.New("delete Favorite Failed")
	}

	authorId := GetVideoAuthorId(videoId)
	ChangeVideoCount(videoId, "favorite_count", -1)
	go ChangeUserCount(userId, "favorite_count", -1)
	go ChangeUserCount(authorId, "favorited_count", -1)

	return nil
}

func GetFavoriteList(userId int64) ([]Video, error) {
	db := common.GetDB()
	var VideoList []Video
	err := db.Joins("left join favorites on favorites.video_id = videos.video_id").Where("favorites.user_id = ?", userId).Find(&VideoList).Error
	if err == gorm.ErrRecordNotFound {
		return VideoList, nil
	} else if err != nil {
		return VideoList, err
	}
	for i, v := range VideoList {
		author, err := GetUserInfo(v.AuthorId)
		if err != nil {
			return VideoList, err
		}
		VideoList[i].Author = author
	}
	return VideoList, nil
}
