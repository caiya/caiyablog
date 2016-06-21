package controllers

import (
	m "caiyablog/models"
	"fmt"
	"path"
	"strconv"
)

const PAGESIZE = 5

type BlogController struct {
	BaseController
}

//博客详情页
func (this *BlogController) BlogInfo() {
	id := path.Base(this.Ctx.Request.RequestURI)
	blog := m.GetBlogById(id)
	if blog.Id == 0 {
		this.TplName = "error_404_2.html"
		return
	}
	this.Data["blog"] = blog
	this.GetAllCategorys()
	this.GetAllArchives(5)
	this.GetBlogs(5)
	this.TplName = "blog.html"
}

//博客列表页
func (this *BlogController) BlogList() {
	blog := new(m.Blog)
	tag := this.GetString(":tag")
	fmt.Println(tag)
	if tag != "" {
		blog.Tag, _ = strconv.Atoi(tag)
	}
	page, _ := this.GetInt64("page")
	if page == 0 {
		page = 1
	}
	blogList, count := m.GetBlogList(blog, page, PAGESIZE)
	for _, v := range blogList {
		v.ToString()
	}
	if count%PAGESIZE == 0 {
		this.Data["count"] = count / PAGESIZE
	} else {
		this.Data["count"] = count/PAGESIZE + 1
	}
	this.Data["blogList"] = blogList
	this.TplName = "blogs.html"
}
