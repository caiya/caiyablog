package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Pass     string `json:"pass"`
	Avatar   string `json:"avatar"`
	Desc     string `json:"desc"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
	Major    string `json:"major"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Nickname string `json:"nickage"`
	Regtime  int    `json:"regtime"`
	Addr     string `json:"addr"`
	Birth    int    `json:"birth"`
	Qq       string `json:"qq"`
	Motto    string `json:"motto"`
}

func (this *User) ToString() {
	fmt.Printf("User: %#v\n", this)
}

func (this *User) TableName() string {
	return "t_user"
}

func init() {
	orm.RegisterModel(new(User))
}

//根据mobile或者id获取用户信息
func (this *User) GetUserByIdOrMobile(name string, arg int64) (user User) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_user")
	if name == "id" {
		qs = qs.Filter("id", arg)
	} else { //默认以手机号查询
		qs = qs.Filter("mobile", arg)
	}
	qs.One(&user)
	return
}
