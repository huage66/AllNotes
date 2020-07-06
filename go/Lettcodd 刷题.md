# Lettcode算法题



##　２．整数反转

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21

**解题思路：反转我们可以使用二分法进行，前面和后面进行交换，涉及到负号我们需要判断一下，为负数我们就把前面索引往前移，最后考虑反转后的int范围**

```go
func reverse(x int) int {
	strArr := strings.Split(strconv.Itoa(x),"")

	for i, j := 0, len(strArr) - 1; i < j;i,j = i + 1, j - 1 {

		if(strings.Compare(strArr[i],"-") == 0){

			i++
		}
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	var result string
	for _, k := range strArr {

		result += k
	}
	s,_ := strconv.Atoi(result)

	if(s < -(1 << 31) || s > (1 << 31) -1 ){

		return 0
	}
	return s

}
```



## ３．回文数


判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

**示例 1:**

```
输入: 121
输出: true
```

**示例 2:**

```
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```

**示例 3:**

```
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

```



**解题思路：也是利用二分法，前面和后面一起索引，查看是否相等**

```go
func isPalindrome(x int) bool {

    arr := strings.Split(strconv.Itoa(x),"")

    for i, j := 0, len(arr) - 1;i < j;i, j = i + 1, j - 1 {

        if strings.Compare(arr[i],"-") == 0 {
            
            return false
        }

        if arr[i] != arr[j] {

            return false
        }
        
    }
    return true
}
```



## 罗马数字转整数

罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.



**解题思路：主要是找规律，罗马数字当后面的一个数字大于前面的一个数字，就需要做减法，否则就是加法，然后结果相加就行。**

```go
func romanToInt(s string) int {

    maps := map[string]int {

        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
    }
    var result int
    str := strings.Split(s,"")
    pre := 0
    for k := len(str) - 1; k >= 0; k-- {

        cur,_ := maps[str[k]]
        if(cur < pre){
            result -=  cur
        }else {
            result += cur
        }
        pre = cur
    }
    return result
}
```

