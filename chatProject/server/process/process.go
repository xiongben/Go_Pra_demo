package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/utils"
	"fmt"
	"io"
	"net"
)

func Processfn(conn net.Conn) {
	defer conn.Close()

	for {
		mes, err := utils.Readpkg(conn)
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
		err = ServerProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		err = ServerProcessLogin(conn, mes)
	case message.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}
