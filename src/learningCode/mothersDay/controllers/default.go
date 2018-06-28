package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "Happy mother's Day!"
	//c.Data["Email"] = "Jeremyyip@github.com"
	//c.TplName = "index.tpl"
	//通过程序获取配置信息
	c.Ctx.WriteString("appname:" + beego.AppConfig.String("appname") +
		"\nhttpport:" + beego.AppConfig.String("httpport") +
		"\nrunmode:" + beego.AppConfig.String("runmode"))

}
