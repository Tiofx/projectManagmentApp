package main

import (
	_ "beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/qor/qor"
)

func init() {
	orm.Debug = true
	log := logs.NewLogger(10000)
	log.SetLogger("console")
	orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1:8088)/web_kp")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
