package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis err = ", err)
		return
	}
	defer conn.Close()
	//redis manage
	_, err = conn.Do("Set", "name", "xiongben")
	if err != nil {
		fmt.Println("manage redis err = ", err)
		return
	}

	res, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("manage redis err = ", err)
		return
	}
	fmt.Println(res)
}
