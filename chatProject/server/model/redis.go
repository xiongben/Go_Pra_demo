package model

import "github.com/garyburd/redigo/redis"

var Pool *redis.Pool

func InitPool(address string) {

	Pool = &redis.Pool{
		TestOnBorrow: nil,
		MaxIdle:      8,
		MaxActive:    0,
		IdleTimeout:  100,
		Wait:         false,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address)
		},
	}
}
