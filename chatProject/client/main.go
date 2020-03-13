package client

import "fmt"

func main() {
	//get user's choice
	var key int
	//judge whether show menu continue
	var loop = true

	for {
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

	} else {

	}
}
