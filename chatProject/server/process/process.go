package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) Processfn() (err error) {
	//defer this.Conn.Close()

	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.Readpkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出。。")
				return err
			} else {
				fmt.Println(err)
				return err
			}
		}
		fmt.Println("mes= ", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
	return
}

func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		up := &UserProcess{Conn: this.Conn}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &UserProcess{Conn: this.Conn}
		err = up.ServerProcessRegister(mes)
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}
