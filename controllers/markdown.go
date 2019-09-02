package controllers

import (
	"blog/define"
	"blog/models"
	"blog/util"
	"fmt"
	"github.com/astaxie/beego/context"
	"strconv"
)

type MarkDownController struct {
	BaseController
}

type MarkContent struct {
	CreateTime 	int64  	`json:"create_time"		form:"create_time"`                   //创建时间
	Title 		string  `json:"title"           form:"title"`
	Content 	string 	`json:"content"         form:"content"`
}


func (this *MarkDownController) AddMarkDown(ctx *context.Context) {

	title := ctx.Request.Form.Get("title")
	content := ctx.Request.Form.Get("content")
	createTime,err := strconv.ParseInt(ctx.Request.Form.Get("created_time"),10,64)

	if err !=nil{
		fmt.Println(err)
		return
	}

	modelMark := models.MarkDown{Title:title,Content:content,CreateTime:createTime}

	if err := modelMark.Add();err !=nil{
		fmt.Println(err)
		return
	}

	util.OutputRes(ctx,&util.Result{
		Code:define.CODE_SUCC,
		Data: &MarkContent{
			Content:content,
			CreateTime:createTime,
			Title:title,
		},
		Msg:"新增成功",
	})

}


