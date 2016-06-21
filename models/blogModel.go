package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Blog struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Type       string `json:"type"`
	Tag        int    `json:"tag"`
	Createtime int    `json:"createtime"`
	Uptime     int    `json:"uptime"`
	Readnum    int    `json:"readnum"`
	Istop      int    `json:"istop"`
	Abstract   string `json:"abstract"`
	Isdelete   int    `json:"isdelete"`
	Userid     int    `json:"userid"`
}

func (this *Blog) ToString() {
	fmt.Printf("blog: %#v\n", this)
}

func (this *Blog) TableName() string {
	return "t_blog"
}

func init() {
	orm.RegisterModel(new(Blog))
}

func GetBlogList(blog *Blog, page int64, pageSize int64) (blogList []*Blog, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_blog")
	var offset int64 = 0
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	if blog.Tag != 0 {
		qs = qs.Filter("tag", blog.Tag)
	}
	qs = qs.Filter("isdelete", 0).Filter("istop", 0)
	qs.Limit(pageSize, offset).OrderBy("-Id").All(&blogList)
	count, _ = qs.Count()
	return
}

//查询首页blog
func GetIndexBlog() (blog Blog) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_blog")
	qs.Filter("istop", 1).One(&blog)
	return
}

//blog查询
func GetBlogs(topCount int, sort string) (blogs []*Blog, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_blog")
	qs.Limit(topCount).Filter("istop", 0).OrderBy(sort).All(&blogs)
	count, _ = qs.Count()
	return
}

//根据id获取博客详情
func GetBlogById(id string) (blog Blog) {
	orm.NewOrm().QueryTable("t_blog").Filter("id", id).Filter("istop", 0).One(&blog)
	return
}
