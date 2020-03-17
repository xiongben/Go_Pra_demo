package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
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
	Code  int    `json:"code"` //500 未注册 200 登录成功
	Error string `json:"error"`
}
