package tool

import "testing"

func TestRegexPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"good case1: number", args{password: "12345678"}, true},
		{"good case2: char", args{password: "abcdefghijk"}, true},
		{"good case2: number and char", args{password: "1234jshsba"}, true},
		{"bad case1: less than the length", args{password: "1"}, false},
		{"bad case2: more than the length", args{password: "12345678901234567"}, false},
		{"bad case3: be not allowed the char", args{password: "#sa;select * from mysql"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegexPassword(tt.args.password); got != tt.want {
				t.Errorf("RegexPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"good case1: number", args{phone: "12345678910"}, true},
		{"bad case2: contain non-number", args{phone: "1234jshsba"}, false},
		{"bad case1: less than the length", args{phone: "1"}, false},
		{"bad case2: more than the length", args{phone: "12345678901234"}, false},
		{"bad case3: be not allowed the char", args{phone: "#sa;select * from mysql"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegexPhone(tt.args.phone); got != tt.want {
				t.Errorf("RegexPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
