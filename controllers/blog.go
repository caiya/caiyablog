package controllers

import (
	m "caiyablog/models"
)

const PAGESIZE = 5

type BlogController struct {
	BaseController
}

//博客详情页
func (this *BlogController) BlogInfo() {
	this.TplName = "blog.html"
}

//博客列表页
func (this *BlogController) BlogList() {
	page, _ := this.GetInt64("page")
	if page == 0 {
		page = 1
	}
	blog := new(m.Blog)
	blogList, count := m.GetBlogList(blog, page, PAGESIZE)
	for _, v := range blogList {
		v.ToString()
	}
	this.Data["blogList"] = blogList
	if count%PAGESIZE == 0 {
		this.Data["count"] = count / PAGESIZE
	} else {
		this.Data["count"] = count/PAGESIZE + 1
	}
	//	this.Data["curr"] = page
	this.TplName = "blogs.html"
}
