package routers

import (
	"github.com/astaxie/beego"
	"mygomvcproject/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/calc/:provinceCode/:cityCode/:amount", &controllers.MainController{}, "get:Calc")

	controllers.InitProvince()
	controllers.InitCityController()
}
