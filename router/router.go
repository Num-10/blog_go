package router

import (
	"blog_go/controller"
	"blog_go/middleware"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	openApi := router.Group("/oo")
	{
		openApi.GET("/", controller.Index)
		openApi.POST("/login", controller.Login)
	}

	authApi := router.Group("/ao")
	authApi.Use(middleware.Verification())
	{
		authApi.GET("/user/:id", controller.User)
		authApi.POST("/tag/save/:id", controller.TagCreate)
		authApi.DELETE("/tag/delete/:id", controller.TagDelete)
		authApi.POST("/article/save/:id", controller.ArticleSave)
	}
}
