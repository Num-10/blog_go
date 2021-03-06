package router

import (
	"blog_go/controller"
	"blog_go/middleware"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router.Static("/api/images", "./runtime/upload/images")
	openApi := router.Group("/api/oo")
	{
		openApi.GET("", controller.Index)
		openApi.GET("/article/:id", controller.SingleArticle)
		openApi.POST("/login", controller.Login)
		openApi.GET("/captcha", controller.GetCaptcha)
		openApi.GET("/captcha/:captchaId", controller.GetCaptchaImage)
		openApi.GET("/verify/:captchaId/:value", controller.VerifyCaptcha)
		openApi.GET("/tag/list", controller.TagList)
		openApi.GET("/time_line", controller.Timeline)
		openApi.GET("/link/list", controller.LinkList)
	}

	authApi := router.Group("/api/ao")
	authApi.Use(middleware.Verification())
	{
		authApi.GET("/user/:id", controller.User)
		authApi.GET("/article/list", controller.ArticleList)
		authApi.GET("/tag/find/:id", controller.TagFind)
		authApi.POST("/tag/save/:id", controller.TagCreate)
		authApi.DELETE("/tag/delete/:id", controller.TagDelete)
		authApi.POST("/article/save/:id", controller.ArticleSave)
		authApi.DELETE("/article/delete/:id", controller.ArticleDelete)
		authApi.GET("/link/find/:id", controller.LinkFind)
		authApi.POST("/link/save/:id", controller.LinkCreate)
		authApi.DELETE("/link/delete/:id", controller.LinkDelete)
		authApi.POST("/upload/image_local", controller.UploadImageLocal)
	}
}
