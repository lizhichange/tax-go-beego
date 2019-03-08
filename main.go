package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "tax-go-beego/db"
	_ "tax-go-beego/routers"
)

func main() {
	orm.Debug = true
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.EnableErrorsShow = true
	beego.SetStaticPath("/swagger", "swagger")
	beego.Run()

}
