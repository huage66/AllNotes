package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	var conn = pool.Get()
	defer conn.Close()

	conn.Do("set", "cat2", "tom")
	line, _ := redis.String(conn.Do("get", "cat1"))

	fmt.Println(line)

}
