package model

import "fmt"

//RManager 房间管理者
var RManager = NewRoomManager(50)

// RoomManager 房间管理者：1表示房间有一个人，表示创建成功，2表示房间有两个人在用，3表示房间已经废弃
type RoomManager struct {
	RoomSize int
	Rooms    []*Room
}

func NewRoomManager(roomSIze int) *RoomManager {
	return &RoomManager{
		RoomSize: roomSIze,
		Rooms:    make([]*Room, 0, roomSIze),
	}
}

func (roomManager *RoomManager) AddRoom(room *Room) {
	if len(roomManager.Rooms) == roomManager.RoomSize {
		fmt.Println("房间不能再创建，已达阙值")
		return
	}
	roomManager.Rooms = append(roomManager.Rooms, room)
}
func (roomManager *RoomManager) GetRoom(roomId int64) *Room {
	for _, room := range roomManager.Rooms {
		if room.RoomId == roomId {
			return room
		}
	}
	return nil
}
