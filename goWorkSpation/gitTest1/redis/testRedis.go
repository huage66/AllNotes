package main

import (
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type hotmap struct {
	hot map[int]int64
}

const (
	MAX_IDLE     = 8
	IDLE_TIMEOUT = 240 * time.Second
)

type MovieList struct {
	currentMovieList []interface{}
}

var m atomic.Value
var mu sync.Mutex

type RedisServer struct {
	Address string
	pool    *redis.Pool
}

var REDIS_NO_SUCH_KEY = errors.New("No such key!")

//func main() {
//	//
//	//server = createRedisServer("localhost:6379")
//	//
//	//
//	//getHotDataEveryday()
//	//
//	//
//	//fmt.Println(currentMovieList)
//}

func stringToArr(str string) []int {

	newArr := str[1 : len(str)-1]
	arr := strings.Split(newArr, " ")
	returnArr := make([]int, 0)
	for _, k := range arr {

		y, _ := strconv.Atoi(k)
		returnArr = append(returnArr, y)
	}

	return returnArr

}

//
func filterHotDay(hot []interface{}, current *[]interface{}, otherSet mapset.Set) {
	count := 0
	//迭代热点数据
	for _, m := range hot {

		if otherSet.Add(m) {
			count++
			*current = append(*current, m)
			if count >= 24 {

				break
			}
		}
	}

}

func getHotDataEveryday() {
	//时间戳
	timeUnix := time.Now().Unix()
	//模拟的热点数据
	hotData := make([]int, 50)

	for i := 200; i < 250; i++ {
		hotData = append(hotData, i)
	}

	//把热点数据放入set集合中
	currentData := mapset.NewSet()
	for _, k := range hotData {

		currentData.Add(k)
	}
	//热点数据,无重复的数据
	hot := currentData.ToSlice()
	//去重之后的热点数据是否少于24个，少就重新再获取
	if len(hot) < 24 {
		//获取之后，再加入，append

	}

	//我们需要查询redis中的数据，如果没有我们就是第一次设置热点数据
	line, err := redis.Int(server.simpleDo("exists", "calendar"))
	if line == 0 {
		//第一次存储，只存储24个热点数据

		line, err := redis.Int(server.simpleDo("zadd", "calendar", timeUnix, hot[0:24]))
		if line == 0 {
			fmt.Println("add hot data failed")
			return
		}
		if err != nil {
			fmt.Println("redis error", err)
			return
		}
		return
	}
	if err != nil {

		fmt.Println("redis get failed", err)
		return
	}

	//不是第一次存储，我们需要进行去重操作，查询该zset集合中的所有数据
	hotmap, err := redis.Strings(server.simpleDo("zrange", "calendar", "0", "-1"))

	if err != nil {
		fmt.Println("redis get failed", err)
	}
	//}else{
	//	fmt.Println(REDIS_NO_SUCH_KEY,err)
	//}

	//把这些数据存入到一个大set中,我们只保存60天数据，之后的删除
	fmt.Println("有多少天的热点数据： ", len(hotmap))
	if len(hotmap) >= 60 {

		//淘汰最小的时间戳，zset集合排序是从小到大
		server.simpleDo("zremrangebyrank", "calendar", 0, 0)
	}

	//把zset中的所有数据存入到set集合中，
	otherSet := mapset.NewSet()
	//数组存入redis后返回的是string

	for _, k := range hotmap {
		newArr := stringToArr(k)
		for _, m := range newArr {

			otherSet.Add(m)
		}

	}
	var current = make([]interface{}, 0)

	//是否有足够的热点数据，没有，我们就再查询之后再过滤
	for len(current) < 24 {
		//过滤数据，过滤掉重复的热点数据
		filterHotDay(hot[:], &current, otherSet)
		//重新加载数据
		for i := 0; i < 50; i++ {
			hot[i] = rand.Intn(10000)
		}
	}

	//存入redis，把值设置给全局变量
	line, e := redis.Int(server.simpleDo("zadd", "calendar", timeUnix, current))
	if line == 0 {
		fmt.Println("add hot data failed")
		return
	}
	if e != nil {

		fmt.Println("redis error", e)
		return
	}
	//设置给全局变量
	//使用automic包中进行原子操作
	mu.Lock()
	defer mu.Unlock()
	m1 := m.Load().(*MovieList)
	fmt.Println("更改前： ", m1.currentMovieList)
	m2 := &MovieList{}
	m2.currentMovieList = current
	fmt.Println("更改后： ", m2.currentMovieList)

	m.Store(m2)

}
func readHotData() {
	for {
		time.Sleep(1 * time.Second)
		m1 := m.Load().(*MovieList)
		fmt.Println("读取热点数据", m1.currentMovieList)
	}
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
