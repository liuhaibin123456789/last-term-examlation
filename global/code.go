package global

/*
	服务端响应码
*/

// ResponseCode 四位数状态码
type ResponseCode int

const (
	CodeSuccess ResponseCode = 1000 + iota
	CodeInvalidParam
	CodeInvalidPathParam
	CodeFailedVote
	CodeUserExist
	CodeUserLoginAlready
	CodeUserNotExist
	CodeUserInvalidPassword
	CodeUserInvalidPhone
	CodeUserInvalidEmail
	CodeUserInvalidName
	CodeUserWrongPassword
	CodeServerBusy //不对外暴露服务器内部错误，统一使用这个返回，错误记录在日志

	CodeEmptyAuth
	CodeInvalidAuth
	CodeWrongToken
	CodeExpiredAccessToken
	CodeExpiredRefreshToken
	CodeNeedLogin

	CodeShouldBindError
	CodeFailedRegister
	CodeFailedLogin

	CodeFailedRefreshToken
	CodeRoomCreateFailed
)

var CodeMap = map[ResponseCode]string{
	CodeSuccess:             "请求成功",
	CodeInvalidParam:        "参数无效，请重新输入",
	CodeUserExist:           "用户已存在，移步登录",
	CodeUserLoginAlready:    "用户别处已登录",
	CodeUserNotExist:        "用户不存在，注册一个吧~",
	CodeUserInvalidPassword: "密码无效，请重新输入",
	CodeUserInvalidPhone:    "手机号无效，请重新输入",
	CodeUserInvalidEmail:    "邮箱格式有误",
	CodeUserInvalidName:     "用户名格式有误",
	CodeUserWrongPassword:   "用户密码错误",
	CodeServerBusy:          "服务繁忙",
	CodeFailedVote:          "投票失败",
	CodeEmptyAuth:           "请求头auth为空",
	CodeInvalidAuth:         "请求头中auth格式有误",
	CodeWrongToken:          "token格式错误",
	CodeNeedLogin:           "需要登录",
	CodeShouldBindError:     "ShouldBind失败",
	CodeFailedRegister:      "注册失败",
	CodeFailedLogin:         "登陆失败",
	CodeExpiredAccessToken:  "access_token时间过期",
	CodeExpiredRefreshToken: "refresh_token时间过期",
	CodeFailedRefreshToken:  "token刷新失败",
	CodeInvalidPathParam:    "无效的路径参数",
	CodeRoomCreateFailed:    "房间创建失败",
}

func (rc ResponseCode) GetMsg() string {
	res, ok := CodeMap[rc]
	if !ok {
		return CodeServerBusy.GetMsg()
	}
	return res
}
