package routers

import (
	"caiyablog/controllers"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/about", &controllers.BaseController{}, "get:About")
	beego.Router("/contact", &controllers.BaseController{}, "get:Contact")
	beego.Router("/message", &controllers.MessageController{}, "post:Message")
	beego.Router("/article/?:tag", &controllers.BlogController{}, "get:BlogList")
	beego.Router("/info/*", &controllers.BlogController{}, "get:BlogInfo")
	beego.ErrorHandler("404", page_not_found)
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("error_404_2.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error_404_2.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
