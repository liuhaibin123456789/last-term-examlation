package tool

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	//获取16位的md5
	newStr := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", newStr)
}
