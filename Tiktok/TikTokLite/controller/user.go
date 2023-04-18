package controller

import (
	"TikTokLite/common"
	. "TikTokLite/message"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) >= 32 || len(password) >= 32 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "输入的账号与密码过长，请重新修改。"},
		})
		return
	}

	RegisterResponse, err := service.UserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, RegisterResponse)
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) >= 32 || len(password) >= 32 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "输入的账号与密码过长，请重新修改。"},
		})
		return
	}
	loginResponse, err := service.UserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, loginResponse)
	}
}

func GetUserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")
	// 这里需要set跟get配合缓存使用，先跳过
	//uid, _ := c.Get("")
	GetUser, err := common.VerifyToken(token)
	//log.Println("0 Success!")
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//log.Println("1 Success!")
	id, _ := strconv.ParseInt(userid, 10, 64)
	mes, err := service.GetUserInfo(GetUser, id)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *mes,
		})
	}
}
