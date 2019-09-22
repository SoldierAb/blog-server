package util

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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
	ctx.Output.SetStatus(http.StatusOK)
	ctx.Output.JSON(target, false, false)
}

func ReadFileOnce(fpath string) (string,error){

	d,err := ioutil.ReadFile(fpath)

	if err !=nil{
		return "",err
	}

	return string(d),nil

}

func SaveContentToFile (content, fpath string) error{
	os.MkdirAll(filepath.Dir(fpath),os.ModePerm)
	fd,err := os.Create(fpath)
	if err !=nil{
		return err
	}
	defer fd.Close()

	_,err = fd.WriteString(content)
	if err !=nil{
		return  err
	}
	return nil
}

func SaveFile(ctx *context.Context,key ,fpath string) error{

	fr ,_ ,err := ctx.Request.FormFile(key)

	if err !=nil{
		if err ==  http.ErrMissingFile{
			return err
		}
		return err
	}

	defer fr.Close()

	//创建上传目录
	os.MkdirAll(filepath.Dir(fpath),os.ModePerm)

	//创建上传文件
	fd,err :=os.Create(fpath)

	if err !=nil{
		return err
	}

	defer fd.Close()

	_,writeErr := io.Copy(fd,fr)

	if writeErr!=nil{
		return writeErr
	}

	return nil

}