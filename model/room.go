package model

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"last-homework/tool"
	"time"
)

const Prepared = "已准备"

// Room 房间
type Room struct {
	RoomId          int64
	Maxsize         int //最多两人
	Status          int //房间的状态：0表示都没有准备，1表示有一个人准备，2表示房间有两个人已准备，可以开始对局
	Clients         map[*Client]bool
	Winner          chan int      //结束,胜利者
	InClient        chan *Client  //进入房间的客户端
	Broadcast       chan []byte   //棋子位置信息进行广播
	SenderClient    *Client       //广播消息发送方
	TimeLimitPerWay time.Duration //等待下棋时间
}

func NewRoom(maxSize int, roomId int64, user *User, conn *websocket.Conn) (*Room, *Client) {
	//房间创建者入房
	client := NewClient(user, conn)

	room := &Room{
		RoomId:          roomId,
		Maxsize:         maxSize,
		Status:          0,
		Clients:         make(map[*Client]bool, maxSize),
		InClient:        make(chan *Client),
		Winner:          make(chan int),
		Broadcast:       make(chan []byte),
		TimeLimitPerWay: time.Second * 30,
	}
	RManager.AddRoom(room)
	tool.SugaredInfo("创建房间成功:", room)
	go room.Start() //开启房间
	room.InClient <- client
	return room, client
}

func (r *Room) Start() {
	defer func() {
		tool.SugaredInfo("关闭房间...")
		close(r.Broadcast)
		close(r.InClient)
		close(r.Winner)
		for c, _ := range r.Clients {
			delete(r.Clients, c)
		}
		//todo 删除数据库
	}()

	for true {
		select {
		case winner := <-r.Winner: //该条chan考虑客户端的退出使用
			//公布结果
			r.SendOther([]byte(MapResult[winner]+" \nGAME OVER!!!"), nil)
			//结束游戏，释放资源
			return
		case c := <-r.InClient:
			r.Clients[c] = true
		case broadcast := <-r.Broadcast:
			//解析广播消息
			msg := &Message{}
			err := json.Unmarshal(broadcast, msg)
			if err != nil {
				tool.SugaredError("json解析失败", err)
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
