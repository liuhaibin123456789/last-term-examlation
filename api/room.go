package api

import (
	"github.com/gin-gonic/gin"
	"last-homework/global"
	"last-homework/model"
	"last-homework/service"
	"last-homework/tool"
	"strconv"
)

func CreateRoom(c *gin.Context) {
	userId := c.GetInt64("user_id")
	//协议升级：长连接
	conn, err := model.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		tool.SugaredError("协议升级失败失败: ", err)
		tool.ResponseError(c, global.CodeServerBusy)
		return
	}
	service.CreateRoom(userId, conn)
}

func EnterRoom(c *gin.Context) {
	userId := c.GetInt64("user_id")
	roomId := c.Query("room_id")
	rId, err := strconv.ParseInt(roomId, 10, 64)
	if err != nil {
		tool.SugaredError("解析失败: ", err)
		tool.ResponseError(c, global.CodeInvalidParam)
		return
	}
	//协议升级：长连接
	conn, err := model.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		tool.SugaredError("协议升级失败: ", err)
		tool.ResponseError(c, global.CodeServerBusy)
		return
	}
	service.EnterRoom(userId, rId, conn)
}
