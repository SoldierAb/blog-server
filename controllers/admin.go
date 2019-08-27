package controllers

import (
	"blog/models"
	"blog/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
)

type AdminController struct {
	BaseController
}

type Admin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	//Isadmin bool
}


func (this *AdminController) Login(ctx *context.Context){
	currentUser := Admin{}

	var code int = 200
	var msg string = "登录成功"


	if err := json.Unmarshal(ctx.Input.RequestBody,&currentUser);err!=nil{
		fmt.Println("Unmarshal error")
		return
	}

	searchAdmin := models.Admin{Username:currentUser.Username}

	fmt.Println(1111,searchAdmin)

	if err := searchAdmin.GetUserByUsername(); err !=nil{
		code = 302
		msg = "查询错误"
	}

	if searchAdmin.Password != currentUser.Password{
		code = 301
		msg = "密码校验失败"
	}

	code  = 200

	util.OutputRes(ctx,&util.Result{
		code,
		currentUser,
		msg,
	})
}
