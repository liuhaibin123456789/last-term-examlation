package api

import (
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegister(t *testing.T) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("user/register", Register)
	go r.Run(":8085")
	url := "http://127.0.0.1:8085/user/register"
	type args struct {
		req *http.Request
		w   *httptest.ResponseRecorder
	}
	tests := []struct {
		name   string
		args   args
		expect int
	}{
		{
			name: "good case1: ",
			args: args{
				req: httptest.NewRequest(http.MethodPost, url, strings.NewReader(`{"phone":"12345678901","password":"shiwer4321"}`)),
				w:   httptest.NewRecorder(),
			},
			expect: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.ServeHTTP(tt.args.w, tt.args.req)
			assert.Equal(t, tt.args.w.Code, tt.expect)
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Login(tt.args.c)
		})
	}
}
