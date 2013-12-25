package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/MartianBlog/models"
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
	this.Data["Title"] = models.BlogInfo["title"]
	this.Data["Subtitle"] = models.BlogInfo["subtitle"]
}
