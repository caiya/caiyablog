package controllers

import (
	m "caiyablog/models"
	"path"
	"regexp"
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

	re, _ := regexp.Compile(`(\d+)年(\d+)月`) //判断是否匹配category类别搜索
	ismatch := re.MatchString(tag)
	if tag != "" && !ismatch {
		blog.Tag, _ = strconv.Atoi(tag) //不满足，表示按照tag搜索
	}
	page, _ := this.GetInt64("page")
	if page == 0 {
		page = 1
	}
	blogList, count := m.GetBlogList(blog, page, PAGESIZE)
	if count%PAGESIZE == 0 {
		this.Data["count"] = count / PAGESIZE
	} else {
		this.Data["count"] = count/PAGESIZE + 1
	}
	this.Data["blogList"] = blogList
	this.TplName = "blogs.html"
}
