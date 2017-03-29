package main

import "fmt"

type F1 float64
type F2 float64


func main() {

	var f1 F1=10.2
	var f2 F2=10.3

	//fmt.Println(f1-f2)//虽然底层都是float64，但是不能直接运算
	fmt.Println(f1-F1(f2))

	/*var int_value int
	int_value=10

	a:=12//简化的方式 常用 相当于 var a in  a=10

	b:="abcdefg"

	_,_,nickname :=name()//匿名变量接收函数的多个返回值
	fmt.Println("nickname is:"+nickname)

	fmt.Println("hello",int_value,a,b)

	*//*const (
		ai=iota//0
		bi=iota//1
		ci=iota//2
	)*//*
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


	fmt.Println(ai,bi,ci,di,ei)*/


	/*a:=new(int)
	b:=new(int)
	fmt.Println(a)//0xc042008290
	fmt.Println(*a)//0
	fmt.Println(a==b)//false*/




}

func name()(firstname string,lastname string,nickname string)  {
	return "Bob","Song","Xiaopang"
}
