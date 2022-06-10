package api

type ReqRegister struct {
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
}
