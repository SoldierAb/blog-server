package controllers

import (
	"blog/define"
	"blog/models"
	"blog/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"path/filepath"
	"strconv"
	"time"
)

type NodeController struct {
	BaseController
}

type NodeRequest struct {
	CategoryId int64 `json:"category_id" form:"category_id"`
	ParentId int64 `json:"parent_id" form:"parent_id"`
	Name string `json:"name" form:"name"`
	Discription string `json:"discription" form:"discription"`
	Content string `json:"content" form:"content"`
}

func (this *NodeController) Add(ctx *context.Context){

	categoryId:= ctx.Input.Query("category_id")
	parentId := ctx.Input.Query("parent_id")
	name := ctx.Input.Query("name")
	discription := ctx.Input.Query("discription")
	content := ctx.Input.Query("content")

	parsedCategoryId,err := strconv.ParseInt(categoryId,10,64)
	if err !=nil{
		this.ServerError(ctx,err)
		return
	}
	parsedParentId , parentErr := strconv.ParseInt(parentId,10,64)
	if parentErr !=nil{
		this.ServerError(ctx,parentErr)
		return
	}

	req := NodeRequest{
		CategoryId:parsedCategoryId,
		ParentId:parsedParentId,
		Name:name,
		Discription:discription,
		Content:content,
	}

	fileStream,has:= ctx.Request.MultipartForm.File[`file`]

	nodeType := define.TYPE_FILE

	ffname := filepath.Join(strconv.Itoa(int(req.CategoryId)),fmt.Sprintf("%s_%d.md",req.Name,time.Now().Unix()))
	fpath:= filepath.Join(beego.AppConfig.String("nodedir"),ffname)

	if has || len(fileStream) != 0{
		err = util.SaveFile(ctx,"file",fpath)
		if err !=nil{
			this.ServerError(ctx,err)
			return
		}

	}else if req.Content != ""{
		err=util.SaveContentToFile(req.Content,fpath)
		if err !=nil{
			this.ServerError(ctx,err)
			return
		}
	}else{
		//this.ServerError(ctx,errors.New(define.CODE_BADREQUEST,define.Msg(define.CODE_BADREQUEST)))
		//return
		nodeType = define.TYPE_DIR
	}


	nodeInstance := models.Node{
		CategoryID:req.CategoryId,
		Type:nodeType,
	}

	conn := models.NewConn()

	insertErr :=nodeInstance.Insert(conn)

	if insertErr !=nil{
		this.ServerError(ctx,err)
		return
	}

	this.Out(ctx,define.CODE_SUCC,nodeInstance)

}
