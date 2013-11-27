package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/MartianBlog/models"
)

type FrontController struct {
	beego.Controller
}

func (this *FrontController) Articles() {
	var articles []models.Article

	if status := this.GetString("status"); status != "" {
		articles = models.ArticlesByStatus(status)
	} else {
		articles = models.AllArticles()
	}

	fmt.Println(articles)
	this.Data["json"] = articles
	this.ServeJson()
}

func (this *FrontController) Article() {
	id := this.GetString("id")

	this.Data["json"] = models.ArticleById(id)
	this.ServeJson()
}
