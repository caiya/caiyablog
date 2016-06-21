package controllers

import (
	m "caiyablog/models"
	"strconv"
	"time"
)

type MessageController struct {
	BaseController
}

func (this *MessageController) Message() {
	message := m.Message{}
	if err := this.ParseForm(&message); err != nil {
		this.resJson(false, err.Error())
		return
	}
	message.Createtime, _ = strconv.Atoi(strconv.FormatInt(time.Now().Unix(), 10))
	message.Ip = this.GetClientIp()
	if message.Ip == "[" {
		message.Ip = "127.0.0.1"
	}
	_, err := m.AddMessage(&message)
	if err != nil {
		this.resJson(false, err.Error())
		return
	}
	this.Redirect("/", 302)
}
