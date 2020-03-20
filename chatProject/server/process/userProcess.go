package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/server/model"
	"awesomeProject1/chatProject/utils"
	"encoding/json"
	"fmt"
	"net"
)

func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	fmt.Println("========")
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
	//fmt.Println(loginMes)
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPass)
	fmt.Println(user)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 500
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
	}
	//if loginMes.UserId == 100 && loginMes.UserPass == 123456 {
	//	loginResMes.Code = 200
	//} else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "用户不存在，请注册后再使用"
	//}

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

func ServerProcessRegister(conn net.Conn, mes *message.Message) (err error) {

}
