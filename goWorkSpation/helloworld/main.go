package main

import (
	"fmt"
	"os"
)

var b bool = false
var a int = 10
var c string = "fjaoefwai"
// int 类型
var demo uint8 = 255
var demo1 int8 = -1
// float 类型
var f1 float32 = 12.164515611
var f2 float32 = -12.164515611
var f3 complex64 = 13.748
// byte 类型  byte相当与uint8 rune相当与int32
var b1 byte = 22
var b2 rune = 3432432;
//uint 32或者64  u什么什么表示只能是整数
var b3 uint =333;
var b4 int = -234;
var b5 uintptr = 212;
func main() {

	// fmt.Println("Hello, world!")
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println("c:"+c)

	// ok()
	// fmt.Println(demo)
	// fmt.Println(f1)
	// fmt.Println(f2)
	// fmt.Println(f3)
	// testVar()
	// testConst()
	// testArtic()
	testCirculateStatement()
	


}

func ok(){
	
	fmt.Println("你好哇，狗子地！");
}

func testVar(){
	//类型推导
	var a,b,c = 1,10,"jdfawoe"
	var ok string
	var test1 int
	var test2 float64
	var test3 bool
	var test4 complex64
	var test5 byte
	

	fmt.Printf("test1=%v test2=%v test3=%v test4=%v test5=%v\n",test1,test2,test3,test4,test5)
	fmt.Println("ok的值是:",ok)
	fmt.Printf("a=%v b=%v c=%v\n",a,b,c)
	var(
		d = 1
		e = "feoa"
		f = true
	)
	fmt.Printf("d=%v e=%v f=%v\n",d,e,f)
	fmt.Printf("f 类型是：f %T \n",f)
	
}

func testConst(){
	// //常量不可改变，声明之后赋值会报错  cannot assign to a
	// const a = 1
	// var b = 1
	// b= 9
	// fmt.Println(a,b)
	// fmt.Printf("a = %T  b = %T\n",a,b)
	
	// const m,f,q = 10,true,"发玩儿哦"
	
	// fmt.Printf("m = %T  f = %T q = %T\n",m,f,q)

	const(

		a = iota
		b 
		c = "fakwoe"
		d = false 
		e 
		f 
		g 

	)

	fmt.Println(a,b,c,d,e,f,g)

}

func testArtic(){

	// var a = 90
	// var b = 40
	// // 0101 1010
	// // 0001 0100
	// // 0001 0000   8
	// // fmt.Printf("a & b = %v",a&b)//8
	//*声明的变量或者常量的地址，&取地址， 然后再次使用*是可以取值
	var ptr *string
	
	var a = "fajoe"
	// 指针类型变量  必须使用&才可以取出声明的地址,  再次使用我们可以用*可以取地址中的值
	ptr = &a
	
	

	fmt.Println("ptr 值为:",ptr)//0xc000010200
	fmt.Println("ptr 值为:",*ptr)//8

	b := 1
	var c int
	c = b 
	c++

	fmt.Println(c)
	


}
// 测试循环语句
func testCirculateStatement(){

	// var a int = 1
	// var b int
	// fmt.Println(a,b)
	// if a == b{

	// 	fmt.Println("真的很帮帮")
	// }else{
	// 	fmt.Println("那就是个贵")
	// }

	// var a = -10
	// var grade string
	// switch  {
	// case a <60 : grade = "没及格"
	// case a >= 60 && a < 80 : grade ="良"
	// case a > 80 : grade="优秀" 
	// case a == 100 : grade="大佬"
	// default: grade = "负分是高手"
	// }
	// fmt.Println(grade)

	// var x interface{}
	// switch i := x.(type){

	// case  nil:
	// 	fmt.Printf("x的类型为： %T\n",i)
		
	// default:
	// 	fmt.Println("无类型")
	// }
	var s, sep string
	for i := 1; i < len(os.Args);i++{

		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

}