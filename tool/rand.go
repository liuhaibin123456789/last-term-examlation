package tool

import (
	"math/rand"
	"time"
)

func RandStr(length int) string {
	if length <= 0 || length > 20 {
		length = 20
	}

	str := "qwertyuiopasdfghjklzxcvbnm"
	var s string
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		s = s + string(str[rand.Intn(len(str))])
	}
	return s
}
