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
	err := json.Unmarshal(this.Ctx.Input.Body(), &article)
	if err != nil {
		panic(err)
	}

	msg := Msg{}
	if ok := models.UpdateArticle(article); ok {
		msg.Code = 200
	} else {
		msg.Code = 500
	}
	fmt.Println(article)
	this.Data["json"] = msg
	this.ServeJson()
}

func (this *AdminController) PostDelete() {
	id := this.GetString("id")
	fmt.Println(id)

	msg := Msg{}
	if ok := models.DeleteArticle(id); ok {
		msg.Code = 200
	} else {
		msg.Code = 500
	}
	this.Data["json"] = msg
	this.ServeJson()
}

func (this *AdminController) UserGet() {
	msg := Msg{}
	email := this.GetSession("email")
	if email == nil {
		msg.Code = 404
	} else {
		user := models.GetUserByEmail(email.(string))
		if user == nil {
			msg.Code = 500
		} else {
			user.Password = ""
			msg.Code = 200
			msg.Data = user
		}
	}

	this.Data["json"] = msg
	this.ServeJson()
}

func (this *AdminController) UserLogin() {
	user := &models.User{}
	err := json.Unmarshal(this.Ctx.Input.Body(), &user)
	if err != nil {
		panic(err)
	}

	msg := Msg{}
	u := models.FindUser(user)
	if u != nil {
		this.SetSession("email", user.Email)
		msg.Code = 200
		u.Password = ""
		msg.Data = u
	} else {
		msg.Code = 403
	}

	this.Data["json"] = msg
	this.ServeJson()
}
