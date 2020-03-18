package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/utils"
	"encoding/json"
	"fmt"
	"net"
)

func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.ummarshal fail err= ", err)
		return
	}
	//response message
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPass == 123456 {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请注册后再使用"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal fail err= ", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal fail err= ", err)
		return
	}
	//发送
	err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("send to client err = ", err)
	}
	return
}
