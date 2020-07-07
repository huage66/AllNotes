package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func CronTask() {
	log.Println("********  *******  *******")
}

func CronTest() {
	log.Println("Starting Cron...")

	c := cron.New()
	c.AddFunc("*/1 * * * * *", CronTask) //
	c.Start()

	t1 := time.NewTimer(time.Second * 5) //
	for {
		select {
		case <-t1.C:
			getHotDataEveryday()
			t1.Reset(time.Second * 5)
		}
	}
}

var server *RedisServer

func main() {
	m.Store(&MovieList{})
	go readHotData()
	server = createRedisServer("localhost:6379")

	CronTest()

}
