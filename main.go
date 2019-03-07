package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "mygomvcproject/db"
	_ "mygomvcproject/routers"
)

func main() {
	orm.Debug = true
	beego.BConfig.EnableErrorsShow = true
	beego.SetStaticPath("/swagger", "swagger")
	beego.Run()

}
