package pkg

import (
	"blog_go/conf"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
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

	/*//写入文件
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Stat(LogPath)
			if os.IsNotExist(err) {
				err = os.MkdirAll(LogPath, os.ModePerm)
				if err != nil {
					fmt.Println("create runtime/log fail：" + err.Error())
				}
			}
			_, err = os.Create(fileName)
			if err != nil {
				fmt.Println("create runtime/log/service.log fail：" + err.Error())
			}
		} else {
			fmt.Println("err", err)
		}
	}*/

	//实例化
	Logger = logrus.New()

	//设置输出
	//Logger.Out = src

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	//设置输出，并进行日志拆分
	Logger.SetOutput(&lumberjack.Logger{
		Filename: fileName,
		MaxSize: conf.LogIni.MaxSize, //Mb
		MaxAge: conf.LogIni.MaxOldLogAge, //days
		MaxBackups: conf.LogIni.MaxOldLogCount,
		LocalTime: true,
	})

	/*//设置 rotatelogs
	logWriter, err := rotatelogs.New(
		//分割后个文件名称
		fileName + ".%Y%m%d.log",

		//生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName),

		//设置最大保存时间（7天）
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//设置日志切割时间间隔（1天）
		rotatelogs.WithRotationTime(24*time.Hour),
		)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	ifHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	Logger.AddHook(ifHook)*/

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
}