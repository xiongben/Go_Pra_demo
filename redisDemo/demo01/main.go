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
	_, err = conn.Do("HMSet", "user02", "name", "Tom", "age", 16)
	if err != nil {
		fmt.Println("manage redis err = ", err)
		return
	}

	res, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("manage redis err = ", err)
		return
	}

	fmt.Println(res)
}
