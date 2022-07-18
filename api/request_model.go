package api

type ReqRegister struct {
	Password string `json:"password" form:"password" bind:"required"`
	Phone    string `json:"phone" form:"phone" bind:"required"`
}
