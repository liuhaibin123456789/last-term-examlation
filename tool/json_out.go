package tool

import (
	"github.com/gin-gonic/gin"
	"last-homework/global"
	"net/http"
)

type ResJson struct {
	Code global.ResponseCode `json:"code,omitempty"`
	Msg  interface{}         `json:"msg,omitempty"`
	Data interface{}         `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code global.ResponseCode) {
	c.JSON(http.StatusOK, ResJson{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, ResJson{
		Code: global.ResponseCode(code),
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResJson{
		Code: global.CodeSuccess,
		Msg:  global.CodeSuccess.GetMsg(),
		Data: data,
	})
}
