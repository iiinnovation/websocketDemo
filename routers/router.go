package routers

import (
	"github.com/astaxie/beego"
	"websocketdemo/controllers"
)

func init ()  {
	beego.Router("/",&controllers.GetWeb{})
	beego.Router("/ws",&controllers.HandellerWebsocket{},"get:Get")
}
