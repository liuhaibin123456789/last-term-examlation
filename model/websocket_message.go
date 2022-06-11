package model

type Message struct {
	//标识消息的唯一性
	Id int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	//聊天室id
	RoomId int64 `json:"room_id" gorm:"type:int;not null"`
	//发表用户的id
	UserId int64 `json:"user_id" gorm:"type:int;not null"`
	//Way 落子下棋的信息.下棋时广播的消息格式：棋子编号+初始坐标+移动后的坐标,传入固定格式json数据，方便解析
	//	{"qi_zi":17,"from_x":0,"from_y":1,"to_x":1,"to_y":0}
	Way []byte `json:"way,omitempty" gorm:"type:varchar(1000);not null"`
	//文字信息的通道
	Info string `json:"info,omitempty" gorm:"type:varchar(1000);not null"`
}

func (m Message) TableName() string {
	return "message"
}

//PlaceChess Message 中的 Message.Way 字段消息格式
type PlaceChess struct {
	//棋子编号,标识棋子
	QiZi  int `json:"qi_zi"`
	FromX int `json:"from_x"`
	FromY int `json:"from_y"`
	ToX   int `json:"to_x"`
	ToY   int `json:"to_y"`
}
