package model

import (
	"blog_go/conf"
	"blog_go/util/e"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var Db *gorm.DB

func SetUp()  {
	db, err := gorm.Open(conf.ModelIni.Connection, conf.ModelIni.Username + ":"+
		conf.ModelIni.Password +"@tcp("+ conf.ModelIni.Host +":"+ conf.ModelIni.Port +")/"+
		conf.ModelIni.Database + conf.ModelIni.Args)
	if err != nil {
		fmt.Println("connect model fail: " + err.Error())
		os.Exit(e.SERVICE_CONNECT_MODEL)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return conf.ModelIni.Prefix + defaultTableName;
	}
	db.SingularTable(true)
	Db = db
}
