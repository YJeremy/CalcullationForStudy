package main

import (
	"beeblog/controllers"
	"beeblog/models" //添加路径
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//数据库初始化
func init() {
	models.RegisterDB() //调用models.go
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.MainController{})
	beego.Run()

}
