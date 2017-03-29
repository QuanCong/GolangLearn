package main

import "fmt"

func main() {

	defer deferFunc1()
	fmt.Println("我定义在deferFunc1之后")
	defer deferFunc2()
	fmt.Println("我定义在deferFunc2之后")

}

func deferFunc1() {
	fmt.Println("我是deferFunc1")
}
func deferFunc2() {
	fmt.Println("我是deferFunc2")
}