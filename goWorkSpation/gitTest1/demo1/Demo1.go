package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var a int8 = 1
var test = []string{"a", "b"}

func printArray(arr *[9]int) { // 引用传递
	for i, v := range arr { //range 关键字
		fmt.Println(i, v)
	}
	arr[0] = 100001
}

func main() {

	//a1 := a
	////a1++
	////fmt.Printf("a=%v  a1=%v\n",a,a1)
	////a2 := &a
	////*a2++
	////fmt.Printf("a=%v  a2=%v",a,*a2)
	//
	//fmt.Println()
	//
	//
	//
	//for _, k := range test {
	//
	//
	//}
	//fmt.Println([]int{1,2,4,5,6})

	file, err := os.Open("hotFile")
	fmt.Println(file.Name())
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	hotDataSrc := make(map[int][]int)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		//把热点元数据存入map集合中
		list := make([]int, 0)
		//读取文件中的字符串会有\n换行作为最后一个结尾  ，所以必须把这个排除掉
		s := strings.Split(line, "\\t")

		index, err := strconv.Atoi(s[0])
		if err != nil {

		}
		val := s[1]
		for _, k := range strings.Split(val[0:len(val)-1], ":") {
			m, err := strconv.Atoi(k)
			if err != nil {

			}
			list = append(list, m)
		}
		hotDataSrc[index] = list
	}

	fmt.Println(hotDataSrc)

}
