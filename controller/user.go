package controller

import (
	"blog_go/util"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	util.Json(c, 200, "user", nil)
}
