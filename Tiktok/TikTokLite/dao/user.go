package dao

import (
	"TikTokLite/common"
	. "TikTokLite/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func UserNameIsExist(username string) error {
	db := common.GetDB()
	user := User{}
	err := db.Where("user_name = ?", username).First(&user).Error
	return err
}

func InsertUser(username string, password string) (User, error) {
	db := common.GetDB()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{
		UserName:        username,
		Password:        string(hash),
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorite:   0,
		FavoritedCount:  0,
		Avatar:          "https://img1.baidu.com/it/u=4072724774,4112808305&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
		BackgroundImage: "https://i0.hdslb.com/bfs/article/6412e2485f8fc179543cecb8de6a726bb525132d.jpg@942w_1565h_progressive.webp",
		Signature:       "",
	}
	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}
	//todo打印一下日志
	return user, nil
}

func GetUserInfo(u interface{}) (User, error) {
	db := common.GetDB()
	user := User{}
	var err error
	switch u.(type) {
	case int64:
		// 后期可以添加缓存cache
		err = db.Where("user_id = ?", u).First(&user).Error
	case string:
		err = db.Where("user_name = ?", u).First(&user).Error
	default:
		err = errors.New("未支持的类型访问")
	}
	return user, err
}

func ChangeUserCount(userid int64, countType string, diff int64) {
	db := common.DB
	var user User
	err := db.Where("user_id = ?", userid).First(&user).Error
	if err != nil {
		log.Fatalln(err)
	}

	switch countType {
	case "favorite_count":
		user.TotalFavorite += diff
	case "favorited_count":
		user.FavoritedCount += diff
	case "follow_count":
		user.FollowCount += diff
	case "follower_count":
		user.FollowerCount += diff
	}

	err = db.Save(&user).Error
	if err != nil {
		log.Fatalln(err)
	}
}
