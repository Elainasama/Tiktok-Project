package service

import (
	"TikTokLite/common"
	"TikTokLite/dao"
	. "TikTokLite/message"
	"TikTokLite/model"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

//var salt = []byte("TikTokLite")

// GenerateHash 对密码进行二次加盐加密
//func GenerateHash(password string) string {
//	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	//salt := generateSalt(10)
//	hash = append(hash, salt...)
//	hash, _ = bcrypt.GenerateFromPassword(hash, bcrypt.DefaultCost)
//	return string(hash)
//}
//
// CompareHashAndPassword 比较密码是否正确
//func CompareHashAndPassword(hash string, password string) bool {
//	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	hashPassword = append(hashPassword, salt...)
//	err := bcrypt.CompareHashAndPassword([]byte(hash), hashPassword)
//	return err == nil
//}

// 生成一个长度为 n 的随机字符串
func generateSalt(n int) []byte {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}

func UserRegister(username string, password string) (*UserLoginResponse, error) {
	err := dao.UserNameIsExist(username)
	if err == nil {
		return nil, errors.New("该用户名已经存在")
	}

	// 在数据库中存入加密后的密码
	user, err := dao.InsertUser(username, password)

	if err != nil {
		return nil, err
	}
	//jwt加密
	token, err := common.GenToken(user.UserId, user.UserName)
	if err != nil {
		return nil, err
	}
	registerResponse := &UserLoginResponse{
		Response: SuccessResponse,
		UserId:   user.UserId,
		Token:    token,
	}
	return registerResponse, nil
}

func UserLogin(username string, password string) (*UserLoginResponse, error) {
	user, err := dao.GetUserInfo(username)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("该用户名未注册")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误！")
	}

	token, err := common.GenToken(user.UserId, username)
	if err != nil {
		return nil, errors.New("获取Token失败")
	}
	return &UserLoginResponse{
		Response: SuccessResponse,
		UserId:   user.UserId,
		Token:    token,
	}, nil
}

func GetUserInfo(GetUserId int64, userid int64) (*User, error) {
	user, err := dao.GetUserInfo(userid)
	if err != nil {
		return nil, err
	}
	mes := messageUserInfo(user)
	if GetUserId != userid {
		mes.IsFollow = dao.IsFollow(GetUserId, user.UserId)
	}
	return mes, nil
}

func GetUserResponse(GetUserId, userid int64) (*UserResponse, error) {
	user, err := GetUserInfo(GetUserId, userid)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		Response: SuccessResponse,
		User:     *user,
	}, nil
}

func messageUserInfo(info model.User) *User {
	return &User{
		Id:              info.UserId,
		Name:            info.UserName,
		FollowCount:     info.FollowCount,
		FollowerCount:   info.FollowerCount,
		IsFollow:        false,
		Avatar:          info.Avatar,
		BackgroundImage: info.BackgroundImage,
		Signature:       info.Signature,
		TotalFavorited:  info.TotalFavorite,
		FavoriteCount:   info.FavoritedCount,
	}
}

func messageUserList(userList []model.User) []User {
	var msgUserList []User
	for _, x := range userList {
		msgUserList = append(msgUserList, *messageUserInfo(x))
	}
	return msgUserList
}
