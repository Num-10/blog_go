package main

import (
	"blog_go/conf"
	"blog_go/middleware"
	"blog_go/model"
	"blog_go/pkg"
	"blog_go/router"
	"github.com/gin-gonic/gin"
)

func init()  {
	conf.SetUp()
	pkg.LogSetUp()
	model.SetUp()
	pkg.RedisSetUp()
}

func main() {
	app := gin.New()

	maxSize := int64(conf.AppIni.MaxMultipartMemory)
	app.MaxMultipartMemory = maxSize << 20 // 3 MiB

	app.Use(middleware.LoggerToFile())
	app.Use(gin.Recovery())

	router.Router(app)

	defer model.Db.Close()
	app.Run(":" + conf.AppIni.Port) // 监听并在 0.0.0.0:8888 上启动服务
}