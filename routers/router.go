package routers

import (
	"github.com/astaxie/beego"
	"tax-go-beego/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/calc/:cityCode/:amount", &controllers.MainController{}, "get:Calc")

	controllers.InitProvince()
	controllers.InitCityController()
}
