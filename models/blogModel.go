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
	Tag        string `json:"tag"`
	Createtime int    `json:"createtime"`
	Uptime     int    `json:"uptime"`
	Readnum    int    `json:"readnum"`
	Istop      int    `json:"istop"`
	Abstract   string `json:"abstract"`
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

func (this *Blog) GetBlogList(blog *Blog, page int64, pageSize int64) (blogList []*Blog, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_blog")
	var total int64 = 0
	if page <= 1 {
		total = page * pageSize
	} else {
		total = (page - 1) * pageSize
	}
	qs.Filter("name", "slene").Limit(page, total).All(&blogList)
	count, _ := qs.Count()
}

//查询首页blog
func (this *Blog) GetIndexBlog() (blog Blog) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_blog")
	qs.Filter("istop", 1).One(&blog)

	return
}

//blog查询
func (this *Blog) GetBlogs(topCount int, sort string) (blogs []*Blog, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_blog")
	qs.Limit(topCount).Filter("istop", 0).OrderBy(sort).All(&blogs)
	count, _ = qs.Count()
	return
}
