package conf

import (
	"blog_go/util"
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

type config_struct struct {
	App
	Model
}

type App struct {
	Name string
	Mode string `ini:"mode"`
	Debug bool
	Port string
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

func init() {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("load app.ini fail: " + err.Error())
		os.Exit(util.SERVICE_ERROR)
	}
	config_load := new(config_struct)
	err = config.MapTo(config_load)
	AppIni = config_load.App
	ModelIni = config_load.Model
	if err != nil {
		fmt.Println("load app.ini fail: " + err.Error())
		os.Exit(util.READ_CONFIG_ERROR)
	}
}
