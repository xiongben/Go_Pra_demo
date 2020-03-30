package model

import (
	"awesomeProject1/chatProject/server/model"
	"net"
)

type CurUser struct {
	Conn net.Conn
	model.User
}
