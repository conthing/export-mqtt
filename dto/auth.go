package dto

type TokenResp struct {
	Token string `json:"token"`
}

type LoginParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
