package routers

import (
	"github.com/astaxie/beego"
	"tax-go-beego/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/calc", &controllers.MainController{}, "post:Calc")

	controllers.InitProvince()
	controllers.InitCityController()
}
