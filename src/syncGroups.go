package main

import (
	"sync"
	"time"
	"fmt"
)

func doSomething(index int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	time.Sleep(1 * time.Second)
	fmt.Println(index, " finished")
}

func main() {
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go doSomething(i, &group)
	}
	group.Wait(); //阻塞等待
	fmt.Println("主 goroutine 退出。。。")
}
