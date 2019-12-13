package controllers

import "github.com/astaxie/beego"

type GetWeb struct {
	beego.Controller
}

func (this *GetWeb)Get()  {
	this.TplName="websocket.html"
}