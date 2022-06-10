package model

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// Upgrade 协议升级
var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}
