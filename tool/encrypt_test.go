// @author cold bin
// @date 2022/7/17

package tool

import "testing"

func TestMD5(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"good case", args{"abcdefg"}, "7ac66c0f148de9519b8bd264312c4d64"},
		{"good case", args{"12345"}, "827ccb0eea8a706c4c34a16891f84e7b"},
		{"good case", args{"#@@#@&*HSJHDJS"}, "c8a70f1c9e204b104431119e4571f625"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.str); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
