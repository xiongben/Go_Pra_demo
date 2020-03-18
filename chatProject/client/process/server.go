package process

import (
	"awesomeProject1/chatProject/utils"
	"fmt"
	"net"
	"os"
)

//显示登陆成功后的界面
func showMenu() {
	fmt.Println("---------XXX login success---------")
	fmt.Println("---------1, 显示在线用户列表---------")
	fmt.Println("---------2， 发送消息---------")
	fmt.Println("---------3， 信息列表---------")
	fmt.Println("---------4， 退出系统---------")
	fmt.Println("请选择1-4 ：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选线不正确")

	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	for {
		mes, err := utils.Readpkg(conn)
		if err != nil {
			fmt.Println("readpkg error = ", err)
			return
		}
		fmt.Println(mes)
	}
}
