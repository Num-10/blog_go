package pkg

import (
	"blog_go/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

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
	SigningKey := []byte(conf.AppIni.SigningKey)
	token, err := tokenClaim.SignedString(SigningKey)
	return token, err
}

func ParseToken(token string) (*CustomClaims, error) {
	customClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.AppIni.SigningKey), nil
	})

	if customClaims != nil {
		if claims, ok := customClaims.Claims.(*CustomClaims); ok && customClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
