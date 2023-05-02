package controller

import (
	"TikTokLite/common"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Register 用户登录
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) >= 32 || len(password) >= 32 {
		ResponseFail(c, "输入的账号与密码过长，请重新修改。")
		return
	}

	RegisterResponse, err := service.UserRegister(username, password)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *RegisterResponse)
}

// Login 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) >= 32 || len(password) >= 32 {
		ResponseFail(c, "输入的账号与密码过长，请重新修改。")
		return
	}

	loginResponse, err := service.UserLogin(username, password)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *loginResponse)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")

	GetUser, err := common.VerifyToken(token)

	if err != nil {
		ResponseFail(c, tokenError)
		return
	}

	id, _ := strconv.ParseInt(userid, 10, 64)
	userResponse, err := service.GetUserResponse(GetUser, id)
	if err != nil {
		ResponseFail(c, err.Error())
		return
	}
	ResponseSuccess(c, *userResponse)
}
