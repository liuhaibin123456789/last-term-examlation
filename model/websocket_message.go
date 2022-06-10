package model

type Message struct {
	//标识消息的唯一性
	Id int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	//聊天室id
	ChatRoomId int64 `json:"chat_room_id" gorm:"type:int;not null"`
	//发表用户的id
	SenderId int64 `json:"sender_id" gorm:"type:int;not null"`
	//内容
	Way string `json:"content" gorm:"type:varchar(1000);not null"`
}

func (m Message) TableName() string {
	return "message"
}
