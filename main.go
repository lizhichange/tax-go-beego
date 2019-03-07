package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	beego.BConfig.EnableErrorsShow = true
	beego.SetStaticPath("/swagger", "swagger")
	beego.Run()

}
