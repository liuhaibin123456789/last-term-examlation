package model

type Message struct {
	//标识消息的唯一性
	Id int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	//聊天室id
	RoomId int64 `json:"room_id" gorm:"type:int;not null"`
	//发表用户的id
	UserId int64 `json:"user_id" gorm:"type:int;not null"`
	//todo 棋子路线
	Way string `json:"content" gorm:"type:varchar(1000);not null"`
}

func (m Message) TableName() string {
	return "message"
}
