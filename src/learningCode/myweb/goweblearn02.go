package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct { //自定义命名一个结构体，继承beego.Controller的内容
	beego.Controller
}

//
func (this *HomeController) Get() {
	this.Ctx.WriteString("hello world")

}

func main() {
	beego.Router("/", &HomeController{}) //注册路由, router  和 controller
	beego.Run()                          // 注册后就开始运行
}
