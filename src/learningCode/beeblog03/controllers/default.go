package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "home.html"

	/*
		c.Data["TrueCond"] = true
		c.Data["FalseCond"] = false
		//直接在模板.tpl里面变成语句用{{}}写编程语句
		//除了<>,其他文字都是直接网页输出的

		type u struct {
			Name string
			Age  int
			Sex  string
		}

		user := &u{
			Name: "Jeremy",
			Age:  29,
			Sex:  "male",
		}

		c.Data["User"] = user
		//打印用 .User.Age
		//嵌套输出
		//with .User
		//end

		num := []int{1, 2, 3, 4, 5, 6}
		c.Data["num"] = num //赋值后还要应用

		//模板变量使用;注意大小写一致;
		//在模板里，调用go的所定义的变量名字，前面要加“.”;
		//在controler 里面处理 ，再渲染出来
		c.Data["TplVar"] = "hey guys"

		//内置模板函数，方便快速处理一些数据，省得在controler里面处理，再渲染出来
		//可以向beego注册自定义的函数
		//内置的html代码转换函数，把安全的html自动转换适合go用的格式
		c.Data["Html"] = "<div>hello beggo html func</div>"

		//pipe line
		c.Data["Pipe"] = "<div>hello beggo html func</div>"
	*/

}
