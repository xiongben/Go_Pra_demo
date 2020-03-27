package process

import (
	"awesomeProject1/chatProject/common/message"
	"awesomeProject1/chatProject/server/model"
)

var onlineUsers map[int]*model.User = make(map[int]*model.User, 10)

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
}
