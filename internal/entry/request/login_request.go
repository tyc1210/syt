package request

type LoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
