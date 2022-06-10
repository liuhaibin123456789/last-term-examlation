package tool

import "regexp"

//RegexPassword 密码：8-16位的字母数字，不包括特殊字符
func RegexPassword(password string) bool {
	regex := `[a-zA-Z0-9]{8,16}`
	c, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}
	return c.MatchString(password)
}

// RegexPhone 手机号：11位数字
func RegexPhone(phone string) bool {
	regex := `[0-9]{11}`
	res, err := regexp.MatchString(regex, phone)
	if err != nil {
		panic(err)
	}
	return res
}
