package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["IsHome"] = true
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Arr"] = []int{1, 2, 2, 3, 3, 4, 4, 55, 6}
	c.TplName = "index.tpl"
}
