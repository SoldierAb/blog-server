package util

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
)


type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
}

func JsonUnmarshal(ctx *context.Context,target interface{}) error{
	return json.Unmarshal(ctx.Input.RequestBody,target)
}

func OutputRes(ctx *context.Context,target interface{}){
	ctx.Output.JSON(target, false, false)
}

