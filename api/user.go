package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"last-homework/global"
	"last-homework/model"
	"last-homework/service"
	"last-homework/tool"
)

func Register(c *gin.Context) {
	user := new(ReqRegister)
	err := c.ShouldBind(user)
	if err != nil {
		//日志记录
		//tool.SugaredWarn("注册api参数有误", err)
		tool.ResponseError(c, global.CodeShouldBindError)
		return
	}
	u := &model.User{
		Password: user.Password,
		Phone:    user.Phone,
	}
	aToken, rToken, err := service.Register(u)
	if err != nil {
		//日志记录
		//tool.SugaredError("service.Register(user)", err)
		tool.ResponseError(c, global.CodeFailedRegister)
		return
	}
	//tool.SugaredInfof("注册用户成功. 手机号：%s,user_id:，%s", user.Phone, u.UserId)
	tool.ResponseSuccess(c, ResRegister{
		RefreshToken: rToken,
		AccessToken:  aToken,
	})

}

func Login(c *gin.Context) {
	user := new(ReqRegister)
	err := c.ShouldBind(user)
	if err != nil {
		//日志记录
		tool.SugaredWarn(
			"登录api参数有误",
			err,
			zap.String("location", "func Login(c *gin.Context)"),
		)
		tool.ResponseError(c, global.CodeShouldBindError)
		return
	}
	u := &model.User{
		Password: user.Password,
		Phone:    user.Phone,
	}
	aToken, rToken, err := service.Login(u)
	if err != nil {
		//日志记录
		tool.SugaredError("service.Login(user)", err)
		tool.ResponseError(c, global.CodeFailedLogin)
		return
	}
	tool.SugaredInfof("登录用户成功. 手机号：%s,user_id:，%s", u.Phone, u.UserId)
	tool.ResponseSuccess(c, ResRegister{
		RefreshToken: rToken,
		AccessToken:  aToken,
	})
}
