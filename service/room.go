package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"last-homework/dao/mysql"
	"last-homework/model"
	"last-homework/tool"
)

func CreateRoom(userId int64, conn *websocket.Conn) {
	user, err := mysql.SelectUser(userId)
	if err != nil {
		return
	}
	//内存
	room, client := model.NewRoom(2, tool.GetId(), user, conn)

	go client.Read(room)
	go client.Write()

	return
}

func EnterRoom(userId, roomId int64, conn *websocket.Conn) {
	user, err := mysql.SelectUser(userId)
	if err != nil {
		return
	}
	client := model.NewClient(user, conn)
	room := model.RManager.GetRoom(roomId)
	if room == nil {
		fmt.Println("没找到room")
		conn.Close()
		return
	}
	room.NotPreparedClient <- client
	go client.Read(room)
	go client.Write()
}
