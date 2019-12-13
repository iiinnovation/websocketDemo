package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"websocketdemo/models"
)

type HandellerWebsocket struct {
	beego.Controller
}

var upgrader =websocket.Upgrader{}
var clients =make(map[*websocket.Conn]bool)

func (this *HandellerWebsocket)Get() {
	//初始化upgrader
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		fmt.Println("初始化upgrader错误", err)
	}
	//延时关闭程序
	defer ws.Close()
	//映射upgrader到clients表里
	clients[ws] = true
	//将json数据反序列送到广播管道
	for {
		//读Message结构体将内容
		var message models.Messages
		err := ws.ReadJSON(&message)
		if err != nil {
			fmt.Println("读取message错误:", err)
			delete(clients, ws)
			break
			/*time.Sleep(time.Second *1)
			message:=models.Messages{Message:"这是向页面发送的数据"+time.Now().Format("2006-01-02 15:04:05")}
			*/
		}
		//将读取内容送到管道
		broadcast <- message
	}

}