package main

import "fmt"

func main() {
	var int_value int
	int_value=10

	a:=12//简化的方式 常用 相当于 var a in  a=10

	b:="abcdefg"

	_,_,nickname :=name()//匿名变量接收函数的多个返回值
	fmt.Println("nickname is:"+nickname)

	fmt.Printf("hello",int_value,a,b)
}

func name()(firstname string,lastname string,nickname string)  {
	return "Bob","Song","Xiaopang"
}
