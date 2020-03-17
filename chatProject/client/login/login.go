package login

import (
	"awesomeProject1/chatProject/common/message"
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
	var bytes []byte
	binary.BigEndian.PutUint32(bytes[0:4], pkglen)
	n, err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("conn.write byte err = ", err)
		return
	}

	fmt.Println("客户端，发送消息长度成功")
	return nil
}
