package router

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine)  {
	openApi := router.Group("/oo"){
		openApi.GET("/")
	}
}
