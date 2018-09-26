package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	autologin := c.Input().Get("autologin") == "on"

	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0

		if autologin {
			maxAge = 1<<31 - 1
		}

		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}

	c.Redirect("/", 301)
	return
}

func CheckAccount(ctx *context.Context) bool {
	cookie, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := cookie.Value

	cookie, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := cookie.Value

	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}
