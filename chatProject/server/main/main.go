package main

import (
	"awesomeProject1/chatProject/server/model"
	"awesomeProject1/chatProject/server/process"
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器在8889端口监听。。。")
	model.InitPool("0.0.0.0:6379")
	initUserDao()
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
		//process.Processfn(conn)
		processcreat(conn)
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(model.Pool)
}

func processcreat(conn net.Conn) {
	processor2 := process.Processor{Conn: conn}
	err := processor2.Processfn()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}
