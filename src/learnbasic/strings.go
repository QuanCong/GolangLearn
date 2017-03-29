package main

import (
	"fmt"
	"strings"
)

func main()  {

	//str1 := `anc what \\cdf` //反引号 原生字符串字面量
	//str2 := "anc what \\" //双引号 可解析的字符串字面量
	//fmt.Println(str1)
	//fmt.Println(str2)

	//fmt.Printf("%t %t\n",true, false)
	//
  	//fmt.Printf("|%b|%9b|%-9b|%09b|% 9b\n",37,37,37,37,37)
  	//fmt.Printf("|%o|%#o|%#8o|\n",37,37,37)
  	//fmt.Printf("|%X|%#x|%#8x|",37,37,37)
	//
	//str_split :="abc-efg-efg-ess-awc"
	//for _,v:=range strings.Split(str_split,"-"){
	//	fmt.Println(v)
	//}
	//fmt.Println(strings.Index(str_split,"eb"))
	//fmt.Println(strings.LastIndex(str_split,"eb"))
	//fmt.Println(strings.IndexAny(str_split,"eb"))//返回b出现在str_split中的位置

	a:="abc"
	b:=strings.Join(a, "bcd")
	fmt.Println(b)


}
