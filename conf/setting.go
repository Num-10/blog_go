package conf

import (
	"blog_go/util/e"
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

type config_struct struct {
	App
	Model
	Redis
	Log
}

type App struct {
	Name string
	Mode string `ini:"mode"`
	Debug bool
	Port string
	MaxMultipartMemory int
	JwtIssuer string
	SigningKey string
}
var AppIni App

type Model struct {
	Connection string
	Host string
	Port string
	Database string
	Username string
	Password string
	Args string
	Prefix string
}
var ModelIni Model

type Redis struct {
	Addr string
	Password string
	Db int
	PoolSize int
}
var RedisIni Redis

type Log struct {
	Runtime string
	LogPath string
	LogFileName string
	MaxSize int
	MaxOldLogAge int
	MaxOldLogCount int
}
var LogIni Log

func SetUp() {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("load app.ini fail: " + err.Error())
		os.Exit(e.SERVICE_ERROR)
	}
	config_load := new(config_struct)
	err = config.MapTo(config_load)
	AppIni = config_load.App
	ModelIni = config_load.Model
	RedisIni = config_load.Redis
	LogIni = config_load.Log
	if err != nil {
		fmt.Println("load app.ini fail: " + err.Error())
		os.Exit(e.READ_CONFIG_ERROR)
	}
}
