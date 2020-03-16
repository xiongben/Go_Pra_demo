package main

import (
	"awesomeProject1/chatProject/client/login"
	"fmt"
)

var userId int
var userPass int

func main() {
	//get user's choice
	var key int
	//judge whether show menu continue
	var loop = true
	for loop {
		fmt.Println("---------welcome to chat room-----------")
		fmt.Println("\t\t\t 1 login in ")
		fmt.Println("\t\t\t 2 register  ")
		fmt.Println("\t\t\t 3 login out  ")
		fmt.Println("\t\t\t please choose one choice ...  ")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("login in chatroom")
			loop = false
		case 2:
			fmt.Println("register")
			loop = false
		case 3:
			fmt.Println("login out")
			loop = false
		default:
			fmt.Println("you put error message")
		}
	}

	if key == 1 {
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%d\n", &userPass)
		err := login.Login(userId, userPass)
		if err != nil {
			fmt.Println("login fail")
		} else {
			fmt.Println("lohin success!")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
}
