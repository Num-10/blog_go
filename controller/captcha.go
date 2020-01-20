package controller

import (
	"blog_go/pkg"
	"blog_go/util/e"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

func GetCaptcha(c *gin.Context)  {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/captcha/" + captchaId + ".png"
	c.JSON(http.StatusOK, captcha)
}

func GetCaptchaImage(c *gin.Context)  {
	pkg.ServeHTTP(c.Writer, c.Request)
}

func VerifyCaptcha(c *gin.Context)  {
	captchaId := c.Param("captchaId")
	value := c.Param("value")
	if captchaId == "" || value == "" {
		e.Json(c, &e.Return{Code:e.PRRAMS_ERROR})
	}
	if captcha.VerifyString(captchaId, value) {
		e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
	} else {
		e.Json(c, &e.Return{Code:e.CAPTCHA_FAIL})
	}
}
