package routers

import (
	"caiyablog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/about", &controllers.BaseController{}, "get:About")
	beego.Router("/contact", &controllers.BaseController{}, "get:Contact")
	beego.Router("/blogs", &controllers.BlogController{}, "get:Blogs")
	beego.Router("/blog", &controllers.BlogController{}, "get:Blog")
}
