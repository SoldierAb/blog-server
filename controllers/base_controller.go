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



func (this *BaseController) Json(ctx *context.Context,data interface{}){
	util.OutputRes(ctx,data)
}


func (this *BaseController) ServerError(ctx *context.Context,err error){
	util.OutputRes(ctx,define.Res(define.CODE_SERVER_ERROR,err))
}

func (this *BaseController) Out(ctx *context.Context,code int,data ...interface{}){
	util.OutputRes(ctx,define.Res(code,data))
}
