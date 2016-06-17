package controllers

type BlogController struct {
	BaseController
}

//博客列表页
func (this *BlogController) Blogs() {
	this.TplName = "blogs.html"
}

//博客详情页
func (this *BlogController) Blog() {
	this.TplName = "blog.html"
}
