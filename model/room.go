package model

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

// Room 房间
type Room struct {
	RoomId            int64
	Maxsize           int //最多两人
	Status            int //房间的状态：0表示都没有准备，1表示有一个人准备，2表示房间有两个人已准备，可以开始对局
	Clients           map[*Client]bool
	PreparedClient    chan *Client  //准备好的房间创建者，默认都未准备
	NotPreparedClient chan *Client  //未准备好的房间创建者
	Broadcast         chan []byte   //棋子位置信息进行广播
	SenderClient      *Client       //广播消息发送方
	TimeLimitPerWay   time.Duration //等待下棋时间
}

func NewRoom(maxSize int, roomId int64, user *User, conn *websocket.Conn) (*Room, *Client) {
	//房间创建者入房
	client := NewClient(user, conn)

	room := &Room{
		RoomId:            roomId,
		Maxsize:           maxSize,
		Status:            0,
		Clients:           make(map[*Client]bool, maxSize),
		NotPreparedClient: make(chan *Client),
		PreparedClient:    make(chan *Client),
		Broadcast:         make(chan []byte),
		TimeLimitPerWay:   time.Second * 30,
	}
	RManager.AddRoom(room)
	fmt.Println("创建房间", room)
	go room.Start() //开启房间
	room.NotPreparedClient <- client
	return room, client
}

func (r *Room) Start() {
	defer func() {
		fmt.Println("关闭房间...")
		close(r.Broadcast)
		close(r.NotPreparedClient)
		close(r.PreparedClient)
		for c, ok := range r.Clients {
			if ok {
				delete(r.Clients, c)
			}
		}
	}()

	for true {
		select {
		case c := <-r.PreparedClient: //该条chan考虑客户端的退出使用
			r.Clients[c] = false
			r.SendOther([]byte("已退出"), c)
		case c := <-r.NotPreparedClient:
			r.Clients[c] = true
			r.sendMe([]byte("请准备"), c)
		case broadcast := <-r.Broadcast:
			//解析广播消息
			msg := &Message{}
			err := json.Unmarshal(broadcast, msg)
			if err != nil {
				fmt.Println(err)
				return
			}
			//提取广播中的棋子路线信息进行广播
			r.SendOther(msg.Way, r.SenderClient)
		}
	}
}

func (r *Room) SendOther(message []byte, ignore *Client) {
	for c := range r.Clients {
		if c != ignore {
			select {
			case c.Message <- message:
			}
		}
	}
}

func (r *Room) sendMe(message []byte, target *Client) {
	select {
	case target.Message <- message:
	}
}
