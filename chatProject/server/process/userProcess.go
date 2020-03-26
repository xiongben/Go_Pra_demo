package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/server/model"
	"awesomeProject1/chatProject/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOthersOnlineUsers(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeToOnlineUsers(userId)
	}
}

func (this *UserProcess) NotifyMeToOnlineUsers(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.marshal err = ", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err = ", err)
		return
	}
	//发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send to client err = ", err)
	}
	return
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
		//把登录成功的用户放到usermgr中
		this.UserId = loginMes.UserId
		userMgr.AddOnlineProcess(this)
		this.NotifyOthersOnlineUsers(loginMes.UserId)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send to client err = ", err)
	}
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.unmarshal fail err = ", err)
		return
	}
	//response message
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var regisResMes message.RegisterResMes
	//fmt.Println(loginMes)
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			regisResMes.Code = 505
			regisResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			regisResMes.Code = 506
			regisResMes.Error = "注册发生未知错误"
		}
	} else {
		regisResMes.Code = 200
	}

	data, err := json.Marshal(regisResMes)
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send to client err = ", err)
	}
	return
}
