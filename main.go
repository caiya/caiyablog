package main

import (
	_ "caiyablog/models"
	_ "caiyablog/routers"

	"github.com/astaxie/beego"
)

const VERSION = "1.0.0"

func main() {
	beego.AppConfig.Set("version", VERSION)
	beego.Run()
}
