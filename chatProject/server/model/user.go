package model

type User struct {
	userId   int    `json:"user_id"`
	UserPass int    `json:"user_pass"`
	UserName string `json:"user_name"`
}
