package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1204)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read message err = ", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("server is listening ,,,")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("wait for the connect")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err = ", err)
		} else {
			fmt.Println("connect =", conn)
			fmt.Println(conn.RemoteAddr())
		}
		go process(conn)
	}
	fmt.Println(listen)
}
