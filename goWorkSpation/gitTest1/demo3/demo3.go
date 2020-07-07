package main

import (
	"fmt"
)

// 声明一个函数类型  type关键字可以声明类型  函数类型也是可以的
type cb func(int) int

func main() {
	//testCallBack(1, callBack)
	//testCallBack(2, func(x int) int {
	//	fmt.Printf("我是回调，x：%d\n", x)
	//	return x
	//})
	var nums = []int{3, 2, 4}
	var target = 6
	fmt.Print(twoSum(nums, target))

}
func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	maps := make(map[int]int, 0)
	for i, k := range nums {

		maps[k] = i
	}

	for i, k := range nums {

		value, ok := maps[target-k]
		if value != i && ok {
			result = append(append(result, i), value)
			break

		}
	}
	return result

}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}
