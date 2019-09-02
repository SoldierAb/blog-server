package main

import (
	"blog/models"
	_ "blog/routers"
	"log"
	"github.com/astaxie/beego"
)


func main() {
	err:=models.InitMysqlDB(&models.Config{
		User:beego.AppConfig.String("user"),
		Password:beego.AppConfig.String("password"),
		Dbname:beego.AppConfig.String("dbname"),
	})

	if err!=nil{
		log.Fatal(err)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
