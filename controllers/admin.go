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

		util.OutputRes(ctx,define.Res(define.CODE_NOT_LOGIN))    //没有登录
		return

	}

	redisToken ,err:= util.GetRedisValue("BLOG-TOKEN")

	if err !=nil{
		util.OutputRes(ctx,define.Res(define.CODE_OVERTIME))    //没有登录
		return
	}

	fmt.Println("redisToken",redisToken,"token",userToken)

	if !strings.EqualFold(userToken,redisToken){
		util.OutputRes(ctx,define.Res(define.CODE_SIGN_IN_OTHER_PLACE))     //token验证不一致
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
		util.OutputRes(ctx,define.Res(define.CODE_USER_NOT_EXISTED))
		return
	}

	if searchAdmin.Password != currentUser.Password{
		util.OutputRes(ctx,define.Res(define.CODE_PASS_WRONG))
		return
	}

	token,err := btoken.CreateToken(&searchAdmin)    //创建token

	if err != nil{
		util.OutputRes(ctx,define.Res(define.CODE_TOKEN_CREATE_ERROR))
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

