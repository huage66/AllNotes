package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"time"
)

func main() {

	CronTest("hotFile")
}
func CronTest(filename string) {
	log.Println("Starting Cron...")

	c := cron.New()
	c.AddFunc("*/1 * * * * *", func() {

	}) //
	c.Start()

	t1 := time.NewTimer(time.Second * 1) //
	for {
		select {
		case <-t1.C:
			ok(filename)
			t1.Reset(time.Second * 5)
		}
	}
}
func ok(filename string) {
	file, err := os.Open(filename)
	if err != nil {

	}
	fmt.Println(file)
	fmt.Println(file.Name())
}
