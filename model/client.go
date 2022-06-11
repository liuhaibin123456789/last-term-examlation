package model

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

//下棋时广播的消息格式：棋子编号+初始坐标+移动后的坐标--》传入固定格式json数据，方便解析{"qi_zi":"棋子编号","from_coordinate":"0 0","to_coordinate":"1 0"}

type Client struct {
	Id       int64               `json:"id"`
	UserName string              `json:"userName"`
	Sex      int8                `json:"sex"`
	Message  chan []byte         `json:"message"`
	CA       *ChessboardAbscissa `json:"ca"` //客户端维持一个棋局
	Conn     *websocket.Conn     `json:"-"`
}

func NewClient(user *User, conn *websocket.Conn) *Client {
	return &Client{
		Id:       user.UserId,
		UserName: user.UserName,
		Sex:      user.Gender,
		Message:  make(chan []byte),
		Conn:     conn,
		CA:       NewChessboardAbscissa(),
	}
}

func (c *Client) Write() {
	fmt.Println(c, " write 开始写入广播")
	defer func() {
		close(c.Message)
		c.Conn.Close()
	}()
	ticker := time.NewTicker(55 * time.Second)
	for true {
		select {
		case message, ok := <-c.Message:
			fmt.Println("write:", string(message))
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					fmt.Println(err)
					return
				}
			}
			//写入客户端
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println(err)
				return
			}
		//心跳,处理ping
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	return
}

//Read 广播消息
func (c *Client) Read(r *Room) {
	fmt.Println(c, " read 开始接收广播")
	defer func() {
		close(c.Message)
		c.Conn.Close()
	}()
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })
	for true {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("read: ", string(message))
		//todo bug 根据消息内容判别是广播还是发给管理端的准备信息.
		if string(message) == "已准备" && r.Status != 2 {
			if r.Status < 3 {
				r.Status++
			}
		} else if r.Status == 2 {
			r.Broadcast <- message
			r.SenderClient = c //广播消息的发送者
		}
	}
}
