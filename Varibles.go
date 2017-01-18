package main

import "fmt"

func main() {
	var int_value int
	int_value=10

	a:=12//简化的方式 常用 相当于 var a in  a=10

	b:="abcdefg"

	_,_,nickname :=name()//匿名变量接收函数的多个返回值
	fmt.Println("nickname is:"+nickname)

	fmt.Println("hello",int_value,a,b)

	/*const (
		ai=iota//0
		bi=iota//1
		ci=iota//2
	)*/
	//上式简化为
	const (
		ai=iota
		bi
		ci
	)


	const (
		di=iota*10// 遇到const关键字 iota 从0开始
		ei=iota*10//10
	)

	const (
		Sunday=iota
		Monday
		Tuesday
		numberofDays //该常量没有导出
	)


	fmt.Println(ai,bi,ci,di,ei)
}

func name()(firstname string,lastname string,nickname string)  {
	return "Bob","Song","Xiaopang"
}
