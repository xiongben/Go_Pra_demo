package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Wait:        false,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	//连接池
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HMSet", "user03", "name", "Jack", "age", 26)
	if err != nil {
		fmt.Println("manage redis err = ", err)
		return
	}

	res, err := redis.Strings(conn.Do("HMGet", "user03", "name", "age"))
	if err != nil {
		fmt.Println("manage redis err = ", err)
		return
	}

	fmt.Println(res)
	pool.Close()

	conn2 := pool.Get() //连接池关闭
	defer conn2.Close()

	res2, err := redis.Strings(conn2.Do("HMGet", "user03", "name", "age"))
	if err != nil {
		fmt.Println("manage redis err = ", err)
		return
	}
	fmt.Println(res2)
}
