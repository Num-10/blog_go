package main

import (
	"blog_go/conf"
	"blog_go/middleware"
	"blog_go/model"
	"blog_go/pkg"
	"blog_go/router"
	"blog_go/util/cron"
	"github.com/gin-gonic/gin"
)

func init()  {
	conf.ConfigSetUp()
	pkg.LogSetUp()
	model.ModelSetUp()
	pkg.RedisSetUp()
	go cron.CronSetup()
}

func main() {
	app := gin.New()

	gin.SetMode(conf.AppIni.Mode)

	maxSize := int64(conf.AppIni.MaxMultipartMemory)
	app.MaxMultipartMemory = maxSize << 20 // 3 MiB

	app.Use(middleware.LoggerToFile())
	app.Use(gin.Recovery())

	router.Router(app)

	defer model.Db.Close()

	app.Run(":" + conf.AppIni.Port) // 监听并在 0.0.0.0:8888 上启动服务
}