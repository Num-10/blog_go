package util

import (
	"blog_go/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var SigningKey = []byte(conf.AppIni.SigningKey)

type CustomClaims struct {
	ID    int `json:"user_id"`
	Name  string `json:"username"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

//创建token
func CreateToken(claims *CustomClaims) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer: conf.AppIni.JwtIssuer,
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(SigningKey)
	return token, err
}
