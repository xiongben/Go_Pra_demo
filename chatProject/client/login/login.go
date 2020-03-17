package login

import (
	"awesomeProject1/chatProject/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
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
	mes, err = readpkg(conn)

	if err != nil {
		fmt.Println("readpkg err = ", err)
		return
	}
	//将mes的data部分反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
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
	return
}

func readpkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn read err = ", err)
		//err = errors.New("read pkg header error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn read fail,err = ",err)
		err = errors.New("read pkg body error")
		return
	}
	//反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
		return
	}
	return
}
