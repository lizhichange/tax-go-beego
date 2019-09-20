package controllers

import "github.com/astaxie/beego"

type VideoController struct {
	beego.Controller
}


func (c *VideoController) Get() {
	c.TplName = "video.html"
}