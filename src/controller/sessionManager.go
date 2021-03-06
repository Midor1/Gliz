package controller

import (
	"config"
	"github.com/garyburd/redigo/redis"
	"github.com/satori/go.uuid"
	"strconv"
)

var RedisConn redis.Conn

func establishRedis() {
	RedisConn, _ = redis.Dial("tcp", config.C.Addresses.RedisAddr)
}

func NewSession(uid int) string {
	establishRedis()
	defer RedisConn.Close()
	u := uuid.NewV4()
	RedisConn.Do("SET", u, strconv.Itoa(uid))
	return u.String()
}

func CheckSession(token string) int {
	establishRedis()
	defer RedisConn.Close()
	uid, err := redis.String(RedisConn.Do("GET", token))
	if err != nil {
		return -1
	}
	id, _ := strconv.Atoi(uid)
	return id
}

func DropSession(token string) int {
	establishRedis()
	defer RedisConn.Close()
	uid, err := redis.String(RedisConn.Do("GET", token))
	if err != nil {
		return -1
	}
	id, err := strconv.Atoi(uid)
	if err != nil {
		return -1
	}
	RedisConn.Do("DEL", token)
	return id
}
