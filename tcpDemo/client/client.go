package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client err = ", err)
		return
	}
	fmt.Println("connect success!", conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err=", err)
		}
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("write err = ", err)
		}
		fmt.Printf("client send the data of %d byte", n)
	}

}
