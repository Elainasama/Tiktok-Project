package dao

import (
	"TikTokLite/common"
	. "TikTokLite/model"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

func InsertMessage(userId int64, toUserid int64, content string) error {
	db := common.GetDB()
	if !isFriend(userId, toUserid) {
		return errors.New("必须和对方成为好友才可以发送信息")
	}
	message := Message{
		ToUserId:   toUserid,
		FromUserId: userId,
		Content:    content,
		CreatTime:  time.Now().Unix(),
	}
	return db.Create(&message).Error
}

func GetMessageList(userId int64, toUserId int64, preMsgTime int64) ([]Message, error) {
	db := common.GetDB()
	var messageList []Message
	err := db.Where("(creat_time > ?) AND ((from_user_id = ? AND to_user_id = ?) "+
		"OR (from_user_id = ? AND to_user_id = ?))",
		preMsgTime, userId, toUserId, toUserId, userId).
		Find(&messageList).Error
	if err == gorm.ErrRecordNotFound {
		return messageList, nil
	}
	return messageList, err
}
