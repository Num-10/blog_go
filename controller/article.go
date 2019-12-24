package controller

import (
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context)  {
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}
