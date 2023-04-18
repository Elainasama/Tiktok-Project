package service

import (
	"TikTokLite/common"
	"TikTokLite/dao"
	. "TikTokLite/message"
	"TikTokLite/model"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(username string, password string) (*UserLoginResponse, error) {
	err := dao.UserNameIsExist(username)
	if err == nil {
		return nil, errors.New("该用户名已经存在")
	}
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
		Response: Response{StatusCode: 0, StatusMsg: "Success!"},
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
		Response: Response{StatusCode: 0, StatusMsg: "Success!"},
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
	//log.Println("2 Success!")
	return mes, nil
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
	var messageUserList []User
	for _, x := range userList {
		messageUserList = append(messageUserList, *messageUserInfo(x))
	}
	return messageUserList
}
