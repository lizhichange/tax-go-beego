package controllers

import (
	"github.com/astaxie/beego"
	"tax-go-beego/models"
)

type CityController struct {
	beego.Controller
}

func InitCityController() {
	beego.Router("/city/getCityByProvinceCode/?:provinceCode", &CityController{}, "get:GetCityByProvinceCode")

}

func (c *CityController) GetCityByProvinceCode() {
	var param = c.Ctx.Input.Param(":provinceCode") //models.GetCityByProvinceCode()
	list, _ := models.GetCityByProvinceCode(param)
	c.Data["json"] = list
	c.ServeJSON()
}
