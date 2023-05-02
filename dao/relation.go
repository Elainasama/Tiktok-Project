package dao

import (
	"TikTokLite/common"
	. "TikTokLite/model"
	"github.com/jinzhu/gorm"
)

// InsertRelation 别忘了修改User数量的变化
func InsertRelation(userId int64, toUserId int64) error {
	db := common.GetDB()
	relation := Relation{
		ToUserId: toUserId,
		FansId:   userId,
	}
	err := db.Create(&relation).Error
	if err != nil {
		return err
	}
	go ChangeUserCount(userId, "follow_count", 1)
	go ChangeUserCount(toUserId, "follower_count", 1)
	return nil
}

// DeleteRelation 别忘了修改User数量的变化
func DeleteRelation(userId int64, toUserId int64) error {
	db := common.GetDB()
	var relation Relation
	err := db.Where("fans_id = ? and to_user_id = ?", userId, toUserId).First(&relation).Error
	if err != nil {
		return err
	}

	err = db.Delete(&relation).Error
	if err != nil {
		return err
	}

	go ChangeUserCount(userId, "follow_count", -1)
	go ChangeUserCount(toUserId, "follower_count", -1)
	return nil
}

func GetFollowUserList(userId int64) ([]User, error) {
	db := common.GetDB()
	var userList []User
	err := db.Joins("left join relations on relations.to_user_id = users.user_id").
		Where("fans_id = ?", userId).
		Find(&userList).Error
	if err == gorm.ErrRecordNotFound {
		return userList, nil
	}
	return userList, err
}

func GetFollowerUserList(userId int64) ([]User, error) {
	db := common.GetDB()
	var userList []User
	err := db.Joins("left join relations on relations.fans_id = users.user_id").
		Where("to_user_id = ?", userId).
		Find(&userList).Error
	if err == gorm.ErrRecordNotFound {
		return userList, nil
	}
	return userList, err
}

func GetFriendList(userId int64) ([]User, error) {
	db := common.GetDB()
	var userList []User
	err := db.Table("users").
		Joins("join relations on relations.fans_id = ? and users.user_id = relations.to_user_id", userId).
		Joins("join relations as r2 on r2.to_user_id = ? and users.user_id = r2.fans_id", userId).
		Find(&userList).Error
	if err == gorm.ErrRecordNotFound {
		return userList, nil
	}
	return userList, err
}

func IsFollow(fansId int64, userId int64) bool {
	db := common.GetDB()
	var relation Relation
	err := db.Where("fans_id = ? and to_user_id = ?", fansId, userId).First(&relation).Error
	return err == nil
}

func isFriend(userId int64, toUserId int64) bool {
	return IsFollow(userId, toUserId) && IsFollow(toUserId, userId)
}
