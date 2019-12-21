package main

import (
	"blog_go/conf"
	"blog_go/model"
	"blog_go/router"
	"github.com/gin-gonic/gin"
)

func init()  {
	conf.Setup()
	model.SetUp()
}

func main() {
	app := gin.New()

	maxSize := int64(conf.AppIni.MaxMultipartMemory)
	app.MaxMultipartMemory = maxSize << 20 // 3 MiB

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	router.Router(app)

	/*router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})*/
	defer model.Db.Close()
	router.Run(":" + conf.AppIni.Port) // 监听并在 0.0.0.0:8888 上启动服务
}