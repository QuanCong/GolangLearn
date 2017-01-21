package main

import "fmt"

func main() {

	value :=23
	pi :=&value
	ppi :=pi
	*ppi++
	fmt.Println(pi)
	fmt.Println(*pi)

	var point_int *int

	point_int=&value
	fmt.Println(point_int)
	v:=*point_int
	fmt.Println(v)



}
