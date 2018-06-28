package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller //beego 架构下 controller文件夹 结构体
}

func (c *LoginController) Get() { //创建函数，Get方法
	c.TplName = "login.html" //指向前端，登录页面
	// 先有login.html 再定义一个 名为c 的结构体指针，指向它，
	//操作c,与之相关连login.html的数据

	isExit := c.Input().Get("exit") == "123" //先判断，然后返回一个bool赋值，获取点击退出
	//c.Ctx.WriteString(isExit)                //输出东西会清空html页面

	//清除不了cookie TO DO!!!
	if isExit { //重新清空设置cookie，maxgae 设置-1 立即删除的意思
		//c.Ctx.WriteString(c.Input().Get("exit")) //可以进入判断isExit 打印这句

		c.Ctx.SetCookie("uname", "123", -1, "/") //右键 查看 页面信息 安全 cookie，这一句有问题，但不知道是怎么设置
		c.Ctx.SetCookie("pwd", "123", -1, "/")
		c.Redirect("/", 301) //重定向到首页
		return               //注意添加return,是因为不需要再渲染login页面
	}

}

func (c *LoginController) Post() {
	//c.Ctx.WriteString(fmt.Sprint(c.Input())) //要配合return
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1

		}

		c.Ctx.SetCookie("uname", uname, maxAge, "/") //右键 查看 页面信息 安全 cookie，
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		//设置了cookie 存去，以及时间;只有正确密码才能存进去;但是下次登录错误了，不过上次cookie时间还存在，cookie不会清空;

	}

	c.Redirect("/", 301) //重定向，配合return;判断右上角的 管理员登录有无 来显示是否登录成功
	return
	//http是无状态协议，每次登录了已经记录在coocki，但是没有告诉网站做再次验证 判断用户是否成功登录的操作

}

//读取cookie ，进行再次验证操作;cookie 成功存进去，但是每次登录页面验证下
func checkAccount(ctx *context.Context) bool { //不需要input 因为是一个cookie的操作
	ck, err := ctx.Request.Cookie("uname") //从请求带上来的cookie
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd

}
