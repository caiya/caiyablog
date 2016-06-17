package controllers

import (
	m "caiyablog/models"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//首页
func (this *BaseController) Get() {

	//类别查询
	this.GetAllCategorys()

	//时间查询
	this.GetAllArchives(7)

	//文章查询
	this.GetBlogs(5)

	//首页单条文章
	b := new(m.Blog)
	blog := b.GetIndexBlog()
	this.Data["blog"] = blog

	this.TplName = "index.html"
}

func (this *BaseController) GetBlogs(count int) {
	blog := new(m.Blog)
	if count <= 0 || count >= 5 {
		count = 5
	}
	blogs, _ := blog.GetBlogs(count, "-createtime")
	this.Data["blogs"] = blogs
}

//查询当前月份向前count月
func (this *BaseController) GetAllArchives(count int) {
	var archives []string
	curMonth := time.Now().Format("2006年01月")
	archives = append(archives, curMonth)
	if count >= 7 || count <= 0 { //最多显示7条
		count = 7
	}
	for i := 0; i < count-1; i++ {
		datetime, _ := time.Parse("2006年01月", archives[i])
		month := datetime.AddDate(0, -1, 0)
		beforeMonth := month.Format("2006年01月")
		archives = append(archives, beforeMonth)
	}
	this.Data["archives"] = archives

}

//类别查询
func (this *BaseController) GetAllCategorys() {
	c := new(m.Category)
	categorys, _ := c.GetAllCategory()
	this.Data["categorys"] = categorys
}

//关于我
func (this *BaseController) About() {
	u := new(m.User)
	curUser := u.GetUserByIdOrMobile("mobile", 18679189528)
	curUser.ToString()
	this.Data["desc"] = curUser.Desc
	this.TplName = "about.html"
}

//联系
func (this *BaseController) Contact() {

	this.TplName = "contact.html"
}
