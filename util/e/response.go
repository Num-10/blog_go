package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Return struct {
	Code int
	Message string
	Data interface{}
}

func Json(c *gin.Context, data *Return)  {
	if data.Data == "" || data.Data == nil || data.Data == 0 {
		data.Data = []string{}
	}
	if data.Message == "" {
		data.Message = GetMsg(data.Code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": data.Code,
		"message": data.Message,
		"data": data.Data,
	})
}

func AbortJson(c *gin.Context, data *Return)  {
	if data.Data == "" || data.Data == nil || data.Data == 0 {
		data.Data = []string{}
	}
	if data.Message == "" {
		data.Message = GetMsg(data.Code)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": data.Code,
		"message": data.Message,
		"data": data.Data,
	})
}
