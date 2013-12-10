package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

type Msg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (this *DefaultController) Get() {
	this.TplNames = "index.html"
}
