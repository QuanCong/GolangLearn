package main

import "fmt"

func main() {
	var i interface{}=100
	var s interface{}=[]string{"1","2"}


	j:=i.(int)
	fmt.Println(i,"->",j)

	if s,ok :=s.([]string);ok{
		fmt.Println(s,"是一个字符串切片")
	}


}
