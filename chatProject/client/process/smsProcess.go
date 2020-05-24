package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

//发送群聊的消息
func (this *SmsProcess) sendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	//序列化
	data, err := json.Marshal(smsMes)
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

	//发送
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send groupMes err = ", err)
	}
	return
}
