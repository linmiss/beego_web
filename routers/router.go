package routers

import (
	"beego_web/controllers"

	"github.com/astaxie/beego"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
}
