package routers

import (
	"caiyablog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/contact", &controllers.MainController{}, "get:Contact")
	beego.Router("/blogs", &controllers.MainController{}, "get:Blogs")
}
