package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	MAX_IDLE     = 8
	IDLE_TIMEOUT = 240 * time.Second
)

type RedisServer struct {
	Address string
	pool    *redis.Pool
}

func createRedisServer(addr string) *RedisServer {

	server := &RedisServer{}
	server.initServer(addr)

	return server

}
func (server *RedisServer) simpleDo(cmd string, args ...interface{}) (interface{}, error) {
	conn := server.getConn()
	defer conn.Close()
	return conn.Do(cmd, args...)
}

func (server *RedisServer) getConn() redis.Conn {

	return server.pool.Get()
}
func (server *RedisServer) getRedisInfo() {

	fmt.Println("addr=%v  pool:%v", server.Address, server.pool.MaxActive)
}
func (server *RedisServer) initServer(addr string) {
	server.Address = addr
	server.pool = newPool(server.Address)
}
func newPool(addr string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     MAX_IDLE,
		IdleTimeout: IDLE_TIMEOUT,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
