package pkg

import (
	"blog_go/conf"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

var Logger *logrus.Logger

func LogSetUp()  {
	Runtime := conf.LogIni.Runtime
	LogPath := conf.LogIni.LogPath
	LogFileName := conf.LogIni.LogFileName

	LogPath = path.Join(Runtime, LogPath)

	//日志文件
	fileName := path.Join(LogPath, LogFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Stat(LogPath)
			if os.IsNotExist(err) {

			}
		} else {
			fmt.Println("err", err)
		}
	}
	//实例化
	Logger = logrus.New()

	//设置输出
	Logger.Out = src

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
}