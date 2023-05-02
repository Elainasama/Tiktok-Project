package dao

import (
	"TikTokLite/common"
	. "TikTokLite/model"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

var UserFilterError = errors.New("user been defend by BloomFilter")

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
		Signature:       "这个人很懒，什么都没有留下",
	}
	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}
	// 缓存存一下新加入的User
	go CacheSetUser(user)
	log.Println("Insert User Success: ", user.UserId, user.UserName)
	return user, nil
}

func GetUserInfo(u interface{}) (User, error) {
	db := common.GetDB()
	user := User{}
	var err error
	switch u := u.(type) {
	case int64:
		// 后期可以添加缓存cache
		user, err = CacheGetUser(u)
		if err == nil {
			return user, nil
		} else if err == redis.Nil {
			err = db.Where("user_id = ?", u).First(&user).Error
		} else {
			return user, err
		}
	case string:
		err = db.Where("user_name = ?", u).First(&user).Error
	default:
		err = errors.New("未支持的类型访问")
	}
	go CacheSetUser(user)
	return user, err
}

func ChangeUserCount(userid int64, countType string, diff int64) {
	strUId := strconv.FormatInt(userid, 10)
	for !common.MutexLock(strUId) {
	}
	db := common.DB
	var user User
	user, err := CacheGetUser(userid)
	if err == redis.Nil {
		err = db.Where("user_id = ?", userid).First(&user).Error
		if err != nil {
			log.Println("Change User Count Error: ", err)
		}
	} else if err != nil {
		log.Println("Change User Count Error: ", err)
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
	CacheSetUser(user)
	err = common.MutexUnlock(strUId)
	if err != nil {
		log.Println("Delete Mutex Failed: ", err)
	}
}

func CacheSetUser(user User) {
	bf := common.GetFilter()
	bf.Add(common.ConvertHash64(strconv.FormatInt(user.UserId, 10)))

	uid := strconv.FormatInt(user.UserId, 10)
	err := common.CacheSet("user"+uid, user)
	if err != nil {
		log.Println("Cache Set User Error: ", err)
	}
	log.Println("Cache Set User Success: ", user.UserId, user.UserName)
}

func CacheGetUser(uid int64) (User, error) {
	bf := common.GetFilter()
	if !bf.Contains(common.ConvertHash64(strconv.FormatInt(uid, 10))) {
		return User{}, UserFilterError
	}
	uidStr := strconv.FormatInt(uid, 10)
	user := User{}
	js, err := common.CacheGet("user" + uidStr)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(js, &user)
	if err != nil {
		log.Println("Cache Get User Error: ", err)
		return user, err
	}
	log.Println("Cache Get User Success: ", user.UserId, user.UserName)
	return user, err
}
