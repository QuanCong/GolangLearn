package main

import "fmt"

func count(channel chan int) {
	fmt.Println("goroutine 启动...")
	fmt.Println(len(channel))
	value := <-channel
	fmt.Println("goroutine 读取到的数据:", value)
}

func write() chan int {
	channel := make(chan int,100)
	go func() {
		channel <- 25
	}()

	return channel
}
func read(channel <-chan int) {
	fmt.Println(<-channel)
}
func changeDirection(channelReadOnly chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for  data := range channelReadOnly {
			out <- data
		}
	}()
	return out
}

func main() {

	chanel := write();
	changeDirection(chanel)
	for v := range chanel {
		fmt.Println(v)
	}
}
