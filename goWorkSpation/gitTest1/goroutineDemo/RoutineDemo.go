package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type MovieList struct {
	currentMovieList int
}

var m atomic.Value
var mu sync.Mutex

func main() {

	m.Store(&MovieList{})
	go updata()
	go read()

	select {}

}
func updata() {

	for {
		time.Sleep(5 * time.Second)

		m1 := m.Load().(*MovieList) // load current value of the data structure
		fmt.Println("更改前： ", m1.currentMovieList)
		m2 := &MovieList{}
		m2.currentMovieList = rand.Intn(100)
		fmt.Println("更改后： ", m2.currentMovieList)

		m.Store(m2)
	}
}

func read() {
	count := 0
	for {
		count++
		time.Sleep(1 * time.Second)
		m1 := m.Load().(*MovieList)

		fmt.Println("读取中.... ", count, m1.currentMovieList)
	}

}
