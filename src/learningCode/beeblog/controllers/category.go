package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")

	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 301) //无论获取与否，都需要重定向，因为需要刷新获取列表
		return

	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}
	c.Data["IsCategory"] = true
	c.TplName = "category.html"
	//正常访问页面

	var err error
	c.Data["Categories"], err = models.GetAllCategories() //c.data已经声明了， 所以单独声明err再一起赋值;返回是两个值，所以要存两个变量

	//要使用一下err
	if err != nil {
		beego.Error(err)
	}

}
