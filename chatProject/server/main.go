package main

import (
	"awesomeProject1/chatProject/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("服务器在8889端口监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen err = ", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来连接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err = ", err)
		}
		//启动一个协程和客户端保持通信
		process(conn)
	}
}

func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

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
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("send to client err = ", err)
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		mes, err := readpkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出。。")
				return
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Println("mes= ", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
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
