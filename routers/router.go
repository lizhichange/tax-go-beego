package routers

import (
	"github.com/astaxie/beego"
	"tax-go-beego/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/calc", &controllers.MainController{}, "post:Calc")

	beego.Router("/getInsuranceByCode/:cityCode", &controllers.MainController{}, "get:GetInsuranceByCode")


	beego.Router("/video", &controllers.VideoController{})

	controllers.InitProvince()
	controllers.InitCityController()
}
