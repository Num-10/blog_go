package main

import (
	"blog_go/conf"
	"blog_go/model"
	"github.com/gin-gonic/gin"
)

func init()  {
	model.SetUp()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	defer model.Db.Close()
	r.Run(":" + conf.AppIni.Port) // 监听并在 0.0.0.0:8888 上启动服务
}