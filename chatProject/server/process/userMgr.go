package process

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{onlineUsers: make(map[int]*UserProcess, 1024)}
}

func (this *UserMgr) AddOnlineProcess(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMgr) DelOnlineProcess(userId int) {
	delete(this.onlineUsers, userId)
}

//查询所有
func (this *UserMgr) GetOnlineProcess() map[int]*UserProcess {
	return this.onlineUsers
}

//查询单个
func (this *UserMgr) FindOnlineProcess(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户id不在线", userId)
		return
	}
	return
}
