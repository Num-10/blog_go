package conf

import (
	"blog_go/util"
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

type config_struct struct {
	App
}

type App struct {
	Name string
	Mode string `ini:"mode"`
	Debug bool
	Port string
}
var AppIni App

func init() {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("load app.ini fail: " + err.Error())
		os.Exit(util.SERVICE_ERROR)
	}
	config_load := new(config_struct)
	err = config.MapTo(config_load)
	AppIni = config_load.App
	fmt.Println(AppIni)
	if err != nil {
		fmt.Println("load app.ini fail: " + err.Error())
		os.Exit(util.READ_CONFIG_ERROR)
	}
}
