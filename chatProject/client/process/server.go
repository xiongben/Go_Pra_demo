package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/utils"
	"encoding/json"
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
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说点什么:")
		fmt.Scanf("%s\n", &content)
		smsProcess.sendGroupMes(content)
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
		//fmt.Println(mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//处理
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//处理
			outputGoupMes(&mes)
		default:
			fmt.Println("服务器返回了未知消息类型")

		}

	}
}
