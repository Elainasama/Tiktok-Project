package dao

import (
	"TikTokLite/common"
	. "TikTokLite/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func InsertComment(userId int64, videoId int64, text string) (Comment, error) {
	db := common.GetDB()
	var comment Comment

	CurrentTime := time.Now()
	month := CurrentTime.Month()
	day := CurrentTime.Day()
	date := fmt.Sprintf("%02d-%02d", month, day)
	comment = Comment{
		UserId:  userId,
		VideoId: videoId,
		Comment: text,
		Time:    date,
	}
	// 对Video数据库的评论数操作
	ChangeVideoCount(videoId, "comment_count", 1)

	return comment, db.Create(&comment).Error
}

func DeleteComment(videoId int64, commentId int64) (Comment, error) {
	db := common.GetDB()
	var comment Comment

	err := db.Where("comment_id = ?", commentId).First(&comment).Error
	if err != nil {
		return comment, err
	}
	err = db.Delete(&comment).Error

	// 对Video数据库的评论数操作
	ChangeVideoCount(videoId, "comment_count", -1)

	return comment, err
}

func GetCommentList(videoId int64) ([]Comment, error) {
	db := common.GetDB()
	var commentList []Comment
	err := db.Where("video_id = ?", videoId).Find(&commentList).Error
	if err == gorm.ErrRecordNotFound {
		return []Comment{}, nil
	} else if err != nil {
		return nil, err
	}
	return commentList, nil
}
