package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.html"
}

func (this *MainController) About() {
	this.TplName = "about.html"
}

func (this *MainController) Contact() {
	this.TplName = "contact.html"
}

func (this *MainController) Blogs() {
	this.TplName = "blogs.html"
}

func (this *MainController) Blog() {
	this.TplName = "blog.html"
}
