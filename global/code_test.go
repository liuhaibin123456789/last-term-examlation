package global

import "testing"

func TestResponseCode_GetMsg(t *testing.T) {
	tests := []struct {
		name string
		rc   ResponseCode
		want string
	}{
		{"good case", CodeUserWrongPassword, "用户密码错误"},
		{"bad case1: not the code", 100000, "服务繁忙"},
		{"bad case2: not the code", 1, "服务繁忙"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rc.GetMsg(); got != tt.want {
				t.Errorf("ResponseCode.GetMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
