package main

import (
	 "blog_go/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println(conf.AppIni.Port)
	fmt.Println(11)
	r.Run(":" + conf.AppIni.Port) // 监听并在 0.0.0.0:8888 上启动服务
}