package main

import (
	"fmt"
	"time"
)

func timeControl(timeout int, chanTimeOut chan bool) {
	time.Sleep(time.Duration(timeout) * time.Second)
	chanTimeOut <- true
}
func dosomething(out chan int) {
	time.Sleep(3 * time.Second)//模拟做一件事
	out <- 1
}
func main() {
	chanDo := make(chan int)
	chanTimeout := make(chan bool)
	go dosomething(chanDo)
	go timeControl(1, chanTimeout)

	select {
	case <-chanDo:
		fmt.Println("任务完成")
	case <-chanTimeout:
		fmt.Println("任务超时")

	}

}
