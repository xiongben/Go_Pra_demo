package model

type User struct {
	UserId     int    `json:"user_id"`
	UserPass   int    `json:"user_pass"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"` //用户状态
}
