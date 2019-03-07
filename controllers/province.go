package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"tax-go-beego/models"
)

type ProvinceController struct {
	beego.Controller
}

func InitProvince() {
	beego.Router("/province/getAllProvince", &ProvinceController{}, "get:GetAllProvince")
}

func (c *ProvinceController) GetAllProvince() {
	var fields []string
	var list, err = models.GetAllProvince(nil, fields, nil, nil, 0, 0)
	fmt.Println(list, err)
	c.Data["json"] = list
	c.ServeJSON()

}
