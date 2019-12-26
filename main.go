package main

import (
	"blog_go/conf"
	"blog_go/model"
	"blog_go/router"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"blog_go/docs"
)

func init()  {
	conf.Setup()
	model.SetUp()
}


// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1/
func main() {
	docs.SwaggerInfo.Title = "Blog APIs"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8888"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := gin.New()

	maxSize := int64(conf.AppIni.MaxMultipartMemory)
	app.MaxMultipartMemory = maxSize << 20 // 3 MiB

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Router(app)

	defer model.Db.Close()
	app.Run(":" + conf.AppIni.Port) // 监听并在 0.0.0.0:8888 上启动服务
}