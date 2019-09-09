package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Criador"] = "Lucas Hokage"
	c.Data["Email"] = "luucas-rb@hotmail.com"
	c.TplName = "index.tpl"
}
