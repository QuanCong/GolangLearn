package main

import (
	"fmt"
)

/**
类型
 */

func main(){
	var byte_value uint16
	byte_value=300

	/*var a int32
	b:= 32 //被自动推导为int类型
	a=32
	fmt.Println(a==b)//编译出错
*/

	var cp complex64
	cp=12+4i
	println(real(cp),imag(cp))

	var str string
	str="hello go lang"
	println(str[0])
	//str[0]='X'//报错，不可以赋值

	//len(str)//取字符串长度
	for index,value :=range str{
		println(index, value)

	}

	array :=[5] int {1,2,3,4,5}

	modify(array)

	for _,value :=range  array{
		println(value)
	}

	fmt.Println(byte_value)


}

func modify(array [5]int)  {
	array[0]=10
	println("函数中打印")
	for _,value :=range  array{
		println(value)
	}
}
