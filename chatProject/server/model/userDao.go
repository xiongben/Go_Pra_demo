package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//在服务器启动后，就初始化一个userDao实例
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	fmt.Println("find user info in redis")
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("jsonunmarshal err = ", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPass int) (User *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	userinfo, err := this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if userinfo.UserPass != userPass {
		err = ERROR_USER_PWD
		return
	}
	return
}
