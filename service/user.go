package service

import (
	"errors"
	"last-homework/dao/mysql"
	"last-homework/model"
	"last-homework/tool"
	"time"
)

var (
	ErrUserPhone         = errors.New("用户手机号格式错误")
	ErrUserPassword      = errors.New("用户密码格式错误")
	ErrUserNotExist      = errors.New("用户手机号不存在")
	ErrUserWrongPassword = errors.New("用户密码错误")
	ErrUserNotUserId     = errors.New("用户id错误")
)

func Register(u *model.User) (aToken, rToken string, err error) {
	//雪花算法生成分布式id
	u.UserId = tool.GetId()
	if u.UserName == "" {
		u.UserName = tool.RandStr(16)
	}
	if !tool.RegexPhone(u.Phone) {
		err = ErrUserPhone
		return
	}
	if !tool.RegexPassword(u.Password) {
		err = ErrUserPassword
		return
	}
	//密码加盐
	u.Password = tool.MD5(u.Password)
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()

	err = mysql.InsertUser(u)
	if err != nil {
		return
	}

	return tool.CreateTwoToken(u.UserId)
}
func Login(u *model.User) (aToken, rToken string, err error) {
	if !tool.RegexPhone(u.Phone) {
		err = ErrUserPhone
		return
	}
	if !tool.RegexPassword(u.Password) {
		err = ErrUserPassword
		return
	}
	//密码加盐
	u.Password = tool.MD5(u.Password)

	user, err := mysql.SelectUserPwd(u.Phone)
	if err != nil {
		err = ErrUserNotExist
		return
	}
	if user.Password != u.Password {
		err = ErrUserWrongPassword
		return
	}
	u.UserId = user.UserId
	return tool.CreateTwoToken(u.UserId)
}
