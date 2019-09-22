package controllers

import (
	"blog/define"
	"blog/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//基类封装
type BaseController struct {
	beego.Controller
}

func (this *BaseController) ServerError(ctx *context.Context,err error){
	util.OutputRes(ctx,define.BaseRes{
		Code:define.CODE_SERVER_ERROR,
		Data:err,
		Msg:define.Msg(define.CODE_SERVER_ERROR),
	})
}


func (this *BaseController) Out(ctx *context.Context,code int,data interface{}){
	util.OutputRes(ctx,define.BaseRes{
		Code:code,
		Data:data,
		Msg:define.Msg(code),
	})
}
