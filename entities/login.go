package entities

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	IsUser bool   `json:"isUser"`
	Token  string `json:"token"`
}
