package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

//发送群聊的消息
func (this *SmsProcess) sendGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.unmarshal err= ", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err= ", err)
		return
	}
	//遍历用户，发送消息
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.sendEachOnlineUser(data, up.Conn)
	}
	return

}

func (this *SmsProcess) sendEachOnlineUser(data []byte, conn net.Conn) (err error) {

	//发送
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send groupMes err = ", err)
	}
	return
}
