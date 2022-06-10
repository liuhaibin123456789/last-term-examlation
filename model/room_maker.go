package model

// RoomMaker 房间创建者表
type RoomMaker struct {
	RoomId int64 `json:"room_id" gorm:"not null"`
	UserId int64 `json:"user_id" gorm:"not null"`
}

func (RoomMaker) TableName() string {
	return "room_maker"
}
