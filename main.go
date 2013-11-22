package main

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/MartianBlog/controllers"
)

func main() {
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/articles", &controllers.FrontController{}, "get:Articles")
	beego.Router("/article", &controllers.FrontController{}, "get:Article")
	beego.Router("/admin/post/new", &controllers.AdminController{}, "post:PostNew")
	beego.Router("/admin/post/update", &controllers.AdminController{}, "post:PostUpdate")
	beego.Run()
}
