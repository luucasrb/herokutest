package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "github.com/saulopinedo/xiaomi-chat"
	c.Data["Email"] = "lucasribeiro@furg.br"
	c.TplName = "index.html"
}
