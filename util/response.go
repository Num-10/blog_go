package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Json(c *gin.Context, code int, message string, data interface{})  {
	if data == "" || data == nil || data == 0 {
		data = []string{}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": message,
		"data": data,
	})
}
