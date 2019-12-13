package controllers

import (
	"fmt"
	"websocketdemo/models"
)

var broadcast =make(chan models.Messages)
func init()  {
	go Sendmsg()
}
func Sendmsg()  {
	for{
		//接收管道信息
		msg:=<-broadcast
		//将接收信息写到页面
		for client :=range clients{
			err:=client.WriteJSON(&msg)
			if err!=nil{
				fmt.Println("写到页面数据错误：",err)
				client.Close()
				delete(clients,client)
			}
		}
	}

}
