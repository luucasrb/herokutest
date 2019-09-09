package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "github.com/luucasrb"
	c.Data["Email"] = "luucas-rb@hotmail.com"
	c.TplName = "index.tpl"
}
