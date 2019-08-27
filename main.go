package main

import (
	"blog/models"
	_ "blog/routers"

	"github.com/astaxie/beego"
)


func main() {
	models.InitMysqlDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
