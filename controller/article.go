package controller

import (
	"blog_go/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context)  {
	util.Json(c, 200, "ok", nil)
}
