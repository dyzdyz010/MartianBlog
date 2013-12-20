package main

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/MartianBlog/controllers"
)

func main() {
	// Front end routers
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/articles", &controllers.FrontController{}, "get:Articles")
	beego.Router("/article", &controllers.FrontController{}, "get:Article")

	// Admin area routers
	beego.Router("/admin/post/new", &controllers.AdminController{}, "post:PostNew")
	beego.Router("/admin/post/update", &controllers.AdminController{}, "post:PostUpdate")
	beego.Router("/admin/post/delete", &controllers.AdminController{}, "delete:PostDelete")
	beego.Router("/admin/user", &controllers.AdminController{}, "get:UserGet")
	beego.Run()
}
