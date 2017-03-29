package main

import (
	"strconv"
	"fmt"
	//"strings"
	"strings"
)

func main(){

	int_v:=20

	//str_bool :="true"

	str_float :="34.56"
	str_int :="   3456"
	str_int=strings.TrimSpace(str_int)
	i1,error:=strconv.ParseInt(str_int,10,16) //3456
	fmt.Println(i1,error)
	i2,error:=strconv.ParseInt(str_int,16,32)//13398 相当于从十六进制值3456转换成十进制的值
	fmt.Println(i2,error)
	i3,error:=strconv.ParseInt(str_int,8,32)//相当于从八进制值3456转换成十进制的值
	fmt.Println(i3,error)
	i4,error:=strconv.ParseInt(str_int,0,32)//自动判断进制
	fmt.Println(i4,error)

	f1,error:=strconv.ParseFloat(str_float,64)
	fmt.Println(f1,error)


	strconv.ParseInt(str_int,int_v,32)//int32
	println(str_int)



	intValue:=23456
	println(&intValue)

	println( strconv.FormatInt(int64(intValue), 8))
	println( strconv.FormatInt(int64(intValue), 10))
	println( strconv.FormatInt(int64(intValue), 2))
	println( strconv.FormatInt(int64(intValue), 16))



}
