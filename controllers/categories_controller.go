package controllers

import (
	"blog/define"
	"blog/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"time"
)

type CategoriesController struct {
	BaseController
}

//文档类别集合
type CategoriesInfo struct {
	*models.Categories
	Dirs models.Nodes
}

func (this CategoriesController) Add(ctx *context.Context){
	cateInstace := models.Categories{}


	if err:= json.Unmarshal(ctx.Input.RequestBody,&cateInstace);err !=nil{
		this.ServerError(ctx,err)
		return
	}

	cateInstace.Created = define.Jsontime(time.Now()).FormatToString()   //创建时间

	db := models.NewConn()

	fmt.Println("cateInstace: ",cateInstace)

	count,err :=cateInstace.CheckName(db)

	if err!=nil{
		this.ServerError(ctx,err)
		return
	}

	if count != 0{
		this.Out(ctx,define.CODE_ALREADY_EXISTED)
		return
	}

	err = cateInstace.Insert(db)

	if err !=nil{
		this.Out(ctx,define.CODE_DATABASE_ERROR,err)
		return
	}

	this.Out(ctx, define.CODE_SUCC)
}

//获取所有类别数据
func (this CategoriesController) GetAll(ctx *context.Context){

	conn := models.NewConn()

	categoriesInstance := models.Categories{}

	list,err:=categoriesInstance.List(conn)

	if err !=nil{
		this.Out(ctx,define.CODE_SERVER_ERROR)
		return
	}

	allCateforiesNodes := make([]CategoriesInfo,len(list))

	for _,val := range list{
		category := CategoriesInfo{
			Categories:val,
		}
		node := models.Node{}
		dirs,err := node.List(conn,category.ID)
		if err !=nil{
			this.Out(ctx,define.CODE_SERVER_ERROR,err)
			return
		}
		category.Dirs = dirs
		allCateforiesNodes = append(allCateforiesNodes,category)
	}

	this.Out(ctx,define.CODE_SUCC,allCateforiesNodes)
}