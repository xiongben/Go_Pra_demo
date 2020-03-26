package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func Login(userId int, userPass int) (err error) {
	//fmt.Println(userId, userPass)
	//return nil

	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net dial err = ", err)
		return
	}

	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPass = userPass

	data, err := json.Marshal(loginMes)
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

	//先发送长度，再发送消息本身
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	_, err = conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("conn.write byte err = ", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write byte(data) err = ", err)
		return
	}

	fmt.Println("客户端，发送消息长度成功", len(data))
	//time.Sleep(10*time.Second)
	//fmt.Println("sleep 10 sec ..")
	mes, err = utils.Readpkg(conn)

	if err != nil {
		fmt.Println("readpkg err = ", err)
		return
	}
	//将mes的data部分反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("userid: ", v)
		}
		go serverProcessMes(conn)
		for {
			showMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func Register(userId int, userPass int, userName string) (err error) {

	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net dial err = ", err)
		return
	}

	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var regisMes message.RegisterMes
	regisMes.User.UserId = userId
	regisMes.User.UserPass = userPass
	regisMes.User.UserName = userName

	data, err := json.Marshal(regisMes)
	if err != nil {
		fmt.Println("json.marshal err = ", err)
		return
	}
	mes.Data = string(data)
	//fmt.Println("==========")
	//fmt.Println(mes)
	//fmt.Println("==========")
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err = ", err)
		return
	}

	//先发送长度，再发送消息本身
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	_, err = conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("conn.write byte err = ", err)
		return
	}
	//fmt.Println(data)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write byte(data) err = ", err)
		return
	}

	fmt.Println("客户端，发送消息长度成功", len(data))
	//time.Sleep(10*time.Second)
	//fmt.Println("sleep 10 sec ..")
	mes, err = utils.Readpkg(conn)

	if err != nil {
		fmt.Println("readpkg err = ", err)
		return
	}
	//将mes的data部分反序列化
	var regisResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &regisResMes)

	if regisResMes.Code == 200 {
		fmt.Println("注册成功")
		go serverProcessMes(conn)
		for {
			showMenu()
		}
	} else {
		fmt.Println(regisResMes.Error)
	}
	return
}
