package main

import (
	_ "testheroku/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

