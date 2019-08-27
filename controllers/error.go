package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404(){

}

func (this * ErrorController) Error501(){

}
