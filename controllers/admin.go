package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/MartianBlog/models"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) PostNew() {
	article := models.Article{}
	err := json.Unmarshal(this.Ctx.Input.Body(), &article)
	if err != nil {
		panic(err)
	}
	fmt.Println(article)
	id := models.AddArticle(article)

	this.Data["json"] = Msg{Code: 200, Data: id}
	this.ServeJson()
}

func (this *AdminController) PostUpdate() {
	article := models.Article{}
	if err := this.ParseForm(&article); err != nil {
		panic(err)
	}
	fmt.Println(article)
}
