package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (this *Category) ToString() {
	fmt.Printf("category: %#v\n", this)
}

func (this *Category) TableName() string {
	return "t_category"
}

func init() {
	orm.RegisterModel(new(Category))
}

func GetAllCategory() (categoryList []*Category, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_category")
	qs.All(&categoryList)
	count, _ = qs.Count()
	return
}
