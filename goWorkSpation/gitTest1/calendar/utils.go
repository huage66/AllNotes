package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"strconv"
	"strings"
	"time"
)

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

func reverArr(arr []int) {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {

		arr[i], arr[j] = arr[j], arr[i]
	}
}

var server *RedisServer

func CronTask() {
	log.Println("********  *******  *******")
}

func CronTest(filename string) {
	log.Println("Starting Cron...")

	c := cron.New()
	c.AddFunc("*/1 * * * * *", CronTask) //
	c.Start()

	t1 := time.NewTimer(time.Second * 1) //
	for {
		select {
		case <-t1.C:
			fmt.Println("定时任务开启==============")
			generateHotData(filename, server.getConn())
			fmt.Println("定时任务结束==============")
			t1.Reset(time.Second * 5)
		}
	}
}
func initAll() {
	//初始化server
	m.Store(&MovieList{})
	server = createRedisServer("127.0.0.1:6379")

	//初始化向redis添加24天的数据
	list := map[int]string{
		1594104669: "35:1",
		1594105000: "36:2",
		1594105001: "37:3",
		1594105002: "39:4",
		1594105003: "41:5",
		1594105004: "40:1",
		1594105005: "42:1",
		1594105006: "43:1",
		1594105007: "44:1",
		1594105008: "45:1",
		1594105009: "46:1",
		1594105010: "47:1",
		1594105011: "48:1",
		1594105012: "49:1",
		1594105013: "50:1",
		1594105014: "51:1",
		1594105015: "52:1",
		1594105016: "53:1",
		1594105017: "54:1",
		1594105018: "55:1",
		1594105019: "56:1",
		1594105020: "57:1",
		1594105021: "58:1",
		1594105022: "59:1"}
	data := map[int][]int{
		36: {1, 2, 3},
		35: {4, 5, 6},
		19: {7, 8, 9},
		94: {10, 11, 12, 13, 14},
		12: {15, 16, 17, 18},
		66: {19, 20, 21, 22},
		89: {23, 24, 25},
		46: {26, 27, 28, 29},
	}

	for i, k := range list {

		server.getConn().Do("zadd", "calendar", i, k)
	}

	//存入备份热点数据
	fmt.Println("覆盖数据")
	server.getConn().Do("del", "srcHotData")
	for i, k := range data {
		server.getConn().Do("hset", "srcHotData", i, k)
	}

	var currentList []int
	for _, k := range list {
		s := strings.Split(k, ":")

		val, _ := strconv.Atoi(s[1])
		currentList = append(currentList, val)
	}
	fmt.Println("初始化过程===")
	updataHotData(currentList)
	getData()
	fmt.Println("初始化完成===")
	//开启任务
	CronTest("hotFile")
	go getHotDataList()
}
