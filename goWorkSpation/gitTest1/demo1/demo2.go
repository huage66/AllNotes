package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//src := []int{1, 2, 3, 4, 5}
	////fmt.Println(src)
	//////temp := append(src[:0],0)
	//////src = append(temp,src)
	////
	//////fmt.Println(src)
	////fmt.Println(time.Now().Unix())//1594104669  1594104683
	//fmt.Println(src[0:0])
	//rs, _ := json.Marshal(src)
	//fmt.Println(rs)
	//
	//json.en

	file, err := os.Open("hotFile")
	if err != nil {

	}
	fmt.Println(file, file.Name())
	buf := bufio.NewReader(file)

	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}

}

func get(data map[int]int) {

	data[1] = 3
}
