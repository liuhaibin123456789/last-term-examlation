package service

import (
	"errors"
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

	//持久化
	err = mysql.CreateRoom(model.RoomMaker{
		RoomId: room.RoomId,
		UserId: userId,
	})
	if err != nil {
		return
	}

	go client.Read(room)
	go client.Write()

	//准备函数: 由管理端接收客户端消息
	room.InClient <- client
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
		tool.SugaredWarn("没找到room")
		conn.Close()
		return
	}
	//表示房间满员
	room.InClient <- client
	go client.Read(room)
	go client.Write()
}

func Search(phone string) (roomId int64, err error) {
	if !tool.RegexPhone(phone) {
		return -1, errors.New("手机格式不对")
	}
	user, err := mysql.SelectUserPwd(phone)
	if err != nil {
		return -1, err
	}
	roomId, err = mysql.SelectRoom(user.UserId)
	return
}
