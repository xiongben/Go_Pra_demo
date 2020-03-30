package process

import (
	"awesomeProject1/chatProject/common/message"
	"encoding/json"
	"fmt"
)

func outputGoupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.unmarshal err=", err.Error())
		return
	}
	info := fmt.Sprintf("userid:\t%d say:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
