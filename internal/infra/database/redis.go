package database

import (
	"github.com/gomodule/redigo/redis"
)

const (
	REDIS_EXPIRE = 60
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	HandleError(err)
	return c
}

func Set(key string, value []byte) error {

	conn := RedisConnect()
	defer conn.Close()

	_, err := conn.Do("SET", key, []byte(value))
	HandleError(err)

	conn.Do("EXPIRE", key, REDIS_EXPIRE)

	return err
}

func Get(key string) ([]byte, error) {

	conn := RedisConnect()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	HandleError(err)

	return data, err
}

func Flush(key string) ([]byte, error) {

	conn := RedisConnect()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("DEL", key))
	HandleError(err)

	return data, err
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
