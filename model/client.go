package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"last-homework/tool"
	"time"
)

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
	tool.SugaredDebug(c, " write 开始写入广播")
	defer func() {
		close(c.Message)
		c.Conn.Close()
	}()
	ticker := time.NewTicker(55 * time.Second)
	for true {
		select {
		case winner, ok := <-c.CA.Winner:
			tool.SugaredDebug(MapResult[winner])
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					tool.SugaredError(err)
					return
				}
			}
			//写入客户端
			if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(MapResult[winner])); err != nil {
				tool.SugaredError(err)
				return
			}
		case message, ok := <-c.Message:
			fmt.Println("write:", string(message))
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					tool.SugaredError(err)
					return
				}
			}
			//写入客户端
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				tool.SugaredError(err)
				return
			}
		//心跳,处理ping
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				tool.SugaredError(err)
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
			tool.SugaredError(err)
			return
		}
		tool.SugaredDebug("read: ", string(message))

		//todo bug 根据消息内容判别是广播还是发给管理端的准备信息.
		for string(message) == Prepared {
			r.Status++
			if r.Status == 2 {
				break
			}
		}
		var msg *Message
		var res bool //是否广播
		//都准备,才能博弈

		//赢了，关闭广播
		if string(message) == MapResult[RedWinner] {
			r.Winner <- RedWinner
		} else if string(message) == MapResult[BlackWinner] {
			r.Winner <- BlackWinner
		} else if string(message) == Prepared {
			//组装message
			msg = &Message{
				Id:     tool.GetId(),
				RoomId: r.RoomId,
				UserId: c.Id,
				Info:   string(message), //客户端之间传递的是棋子信息
			}
		} else {
			//组装message
			msg = &Message{
				Id:     tool.GetId(),
				RoomId: r.RoomId,
				UserId: c.Id,
				Way:    message, //客户端之间传递的是棋子信息
			}
			//解析message的way,操作自己的棋盘
			chess := &PlaceChess{}
			err = json.Unmarshal(message, chess)
			if err != nil {
				tool.SugaredError("json解析失败", err)
				return
			}
			chessRule := c.CA.ChessRule(chess.QiZi, chess.FromY, chess.FromX, chess.ToY, chess.ToX)
			//操作自己的棋盘
			if !chessRule {
				//棋子走错，重新重新走
				r.sendMe([]byte("路线错误，请重新发送走棋的路线！"), c)
				res = false
			} else { //棋子路线正确，广播给另一个客户端
				//打印地图至客户端
				c.CA.PrintChessboard()
				res = true
			}

			if res {
				bytes1, err := json.Marshal(msg)
				if err != nil {
					tool.SugaredError("json编码失败", err)
					return
				}
				//去除json数据空格和下划线
				bytes1 = bytes.TrimSpace(bytes.Replace(bytes1, []byte("\n"), []byte(" "), -1))
				r.Broadcast <- bytes1
				r.SenderClient = c //广播消息的发送者
			}
		}
	}
}
