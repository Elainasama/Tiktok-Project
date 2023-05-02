package common

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTClaims struct {
	Id   int64
	Name string
	jwt.RegisteredClaims
}

var (
	secret = []byte("TikTokLite")
)

// jwt生成
func GenToken(userId int64, userName string) (string, error) {
	claim := JWTClaims{
		Id:   userId,
		Name: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "server",                                           //签发人
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	SignedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return SignedToken, nil
}

// ParseRegisteredClaims 解析jwt
func ParseRegisteredClaims(tokenString string) (*JWTClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil { // 解析token失败
		return nil, err
	}
	if claim, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claim, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

// VerifyToken 验证jwt如果成功就返回userid
func VerifyToken(tokenString string) (int64, error) {
	if tokenString == "" {
		return int64(0), nil
	}
	claim, err := ParseRegisteredClaims(tokenString)
	if err != nil {
		return 0, err
	}
	return claim.Id, nil
}
