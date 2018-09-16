package main

import (
	"beego_web/models"
	"beego_web/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	routers.Init()
	beego.Run()
}
