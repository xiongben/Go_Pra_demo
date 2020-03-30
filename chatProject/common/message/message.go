package message

import "awesomeProject1/chatProject/server/model"

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"user_id"`
	UserPass int    `json:"user_pass"`
	UserName string `json:"user_name"`
}

type LoginResMes struct {
	Code    int    `json:"code"` //500 未注册 200 登录成功
	UserIds []int  `json:"user_ids"`
	Error   string `json:"error"`
}

type RegisterMes struct {
	User model.User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//上线通知用户状态变化
type NotifyUserStatusMes struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

//smsMes
type SmsMes struct {
	Content string `json:"content"`
	model.User
}
