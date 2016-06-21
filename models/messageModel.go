package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"enail" form:"email"`
	Subject    string `json:"subject" form:"subject"`
	Content    string `json:"content" form:"content"`
	Createtime int    `json:"createtime"`
	Ip         string `json:"ip"`
}

func init() {
	orm.RegisterModel(new(Message))
}

func (this *Message) ToString() {
	fmt.Printf("message:%#v\n", this)
}

func (this *Message) TableName() string {
	return "t_message"
}

//留言
func AddMessage(message *Message) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(message)
	return id, err
}
