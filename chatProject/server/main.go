package main

import (
	"fmt"
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

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn read err = ", err)
			return
		}
		fmt.Println("读到的buf=", buf)
	}
}
