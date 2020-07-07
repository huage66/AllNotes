package main

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){

	fmt.Println(reverse(121142))



}

func reverse(x int) int {

	src := strconv.Itoa(x)
	srcArr := strings.Split(src,"")

	for i, j := 0 ,len(srcArr) - 1;i < j;i, j = i + 1,j - 1{

		srcArr[i] ,srcArr[j] = srcArr[j], srcArr[i]
	}
	var result string
	for _, k := range srcArr {

		result += k
	}
	s,_ := strconv.Atoi(result)
	return s
}

func longestCommonPrefix(strs []string) string {

	res := ""
	first := strs[0]
	for i,j := 1, 0; i < len(strs); i++ {


	}


}

