package main

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/MartianBlog/controllers"
)

func main() {
	beego.Router("/admin", &controllers.AdminController{})
	beego.Run()
}
