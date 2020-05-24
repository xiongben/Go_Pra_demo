package process

import (
	model2 "awesomeProject1/chatProject/client/model"
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/server/model"
	"fmt"
)

var onlineUsers map[int]*model.User = make(map[int]*model.User, 10)
var CurUser model2.CurUser //用户登录成功后进行初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	for id, _ := range onlineUsers {
		fmt.Println("用户id： ", id)
	}
}

//处理返回的信息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &model.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
