package controller

import (
	"blog_go/model"
	"blog_go/util"
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary 获取用户基础信息
// @Description 根据user_id获取用户基础信息
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "user_id"
// @Success 200 {object} util.CustomClaims
// @Header 200 {string} Token "qwerty"
// @Router /ao/user/{id} [get]
func User(c *gin.Context) {
	user, _ := c.Get("login_user")
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data:user})
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
