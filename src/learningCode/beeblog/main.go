package main

import (
	"beeblog/controllers"
	"beeblog/models" //添加路径
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//数据库初始化
func init() { //注意函数名必须是 init
	models.RegisterDB() //调用models.go
}

func main() {
	//开启 ORM调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)

	//注册 beego路由
	beego.Router("/", &controllers.MainController{})             //注册的是首页的路由
	beego.Router("/login", &controllers.LoginController{})       // 注册登录页面04
	beego.Router("/category", &controllers.CategoryController{}) // 注册登录页面04

	//启动 beego
	beego.Run()

}
