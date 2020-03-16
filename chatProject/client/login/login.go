package login

import "fmt"

func Login(userId int, userPass int) (err error) {
	fmt.Println(userId, userPass)
	return nil
}
