package controller

import (
	"blog_go/model"
	"blog_go/util"
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	e.Json(c, &e.Return{Code:e.TOKEN_IN_VAIN})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		e.Json(c, &e.Return{Code:e.LOGIN_PARAM_EMPTY})
		return
	}
	where := model.User{
		Username: username,
		Password: password,
	}
	user := new(model.User)
	user.GetUser(&where)
	if user.ID <= 0 {
		e.Json(c, &e.Return{Code:e.LOGIN_PARAM_ERROR})
		return
	}
	claim := util.CustomClaims{
		ID: user.ID,
		Name: user.Username,
	}
	token, err := util.CreateToken(&claim)
	if err != nil {
		e.Json(c, &e.Return{Code:e.TOKEN_CREATE_FAIL})
		return
	}
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data:map[string]string{"token": token}})
}
