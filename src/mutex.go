package main

import (
	"sync"
	"fmt"
	"runtime"
)

var mu sync.Once
var resource map[string]int

func Init() {
	resource = make(map[string]int, 20)
	resource["1"] = 1
}
func getValue() int {
	mu.Do(Init)
	return resource["1"]
}

func main() {
	runtime.GOMAXPROCS(2)
	for{
		go fmt.Print(0)
		fmt.Print(1)
	}


}
