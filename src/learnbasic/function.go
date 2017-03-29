package main

import (
	"fmt"
)

func myprint(args...interface{}) {
	for _, value := range args {
		switch value.(type) {
		case int:
			fmt.Println(value, "is a int type")
		case string:
			fmt.Println(value, " is a string type")
		case float32:
			fmt.Println(value, " is a float32 type")
		default:
			fmt.Println(value, "unkown type")

		}
	}



}
func main() {
	a := func() {
		println("这是闭包a")
	}
	a()

}
