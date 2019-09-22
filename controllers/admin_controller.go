package controllers

import (
	"blog/define"
	"blog/models"
	btoken "blog/token"
	"blog/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"log"
	"strings"
)

type AdminController struct {
	BaseController
}

type Admin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	//Isadmin bool
}


//token校验
func (this *AdminController) Authentication(ctx *context.Context){

	var userToken string


	if userToken = ctx.Request.Header.Get("token");userToken == ""{

		//util.OutputRes(ctx,define.Res(define.CODE_NOT_LOGIN))    //没有登录
		this.Out(ctx,define.CODE_NOT_LOGIN,nil)
		return

	}

	redisToken ,err:= util.GetRedisValue("BLOG-TOKEN")

	if err !=nil{
		this.Out(ctx,define.CODE_OVERTIME,nil)    //没有登录
		return
	}

	fmt.Println("redisToken",redisToken,"token",userToken)

	if !strings.EqualFold(userToken,redisToken){
		this.Out(ctx,define.CODE_SIGN_IN_OTHER_PLACE,nil)  //token验证不一致
		return
	}

	username,password,err :=  btoken.AuthToken(userToken)

	if(err!=nil){
		log.Fatal("authtoken error :  ",err)
		return
	}

	fmt.Println("解析结果：  ",username,password)

	return
}


func (this *AdminController) Login(ctx *context.Context){

	currentUser := Admin{}

	if err := json.Unmarshal(ctx.Input.RequestBody,&currentUser);err!=nil{
		fmt.Println("Unmarshal error")
		return
	}

	searchAdmin := models.Admin{Username:currentUser.Username}

	if err := searchAdmin.GetUserByUsername(); err !=nil{
		this.Out(ctx,define.CODE_USER_NOT_EXISTED,nil)
		return
	}

	if searchAdmin.Password != currentUser.Password{
		this.Out(ctx,define.CODE_PASS_WRONG,nil)
		return
	}

	token,err := btoken.CreateToken(&searchAdmin)    //创建token

	if err != nil{
		this.Out(ctx,define.CODE_TOKEN_CREATE_ERROR,nil)
		return
	}

	util.OutputRes(ctx,define.Res(define.CODE_SUCC,struct {
		Username  string `json:"username"`
		Token string `json:"token"`
	}{
		Username:currentUser.Username,
		Token:token,
	}))

}

