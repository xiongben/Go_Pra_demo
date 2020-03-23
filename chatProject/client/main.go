package main

import (
	"awesomeProject1/chatProject/client/process"
	"fmt"
)

var userId int
var userPass int
var userName string

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
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%d\n", &userPass)
			process.Login(userId, userPass)
			//loop = false
		case 2:
			fmt.Println("register")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入设置密码")
			fmt.Scanf("%d\n", &userPass)
			fmt.Println("请输入用户名称")
			fmt.Scanf("%v\n", &userName)
			process.Register(userId, userPass, userName)
			//loop = false
		case 3:
			fmt.Println("login out")
			//loop = false
		default:
			fmt.Println("you put error message")
		}
	}

	if key == 1 {

		//if err != nil {
		//	fmt.Println("login fail")
		//} else {
		//	fmt.Println("login success!")
		//}
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
}
