package main

import (
	"fmt"
)

func readValue(chanel chan int, done chan bool) {
	for value := range chanel {
		fmt.Println("receive：", value)
	}
	done <- true

}

func sendValue(channel chan int) {
	for i := 0; i < 3; i++ {
		channel <- i
	}
	close(channel)//若不关闭

}

func main() {

	channel := make(chan int,10)
	done := make(chan bool)
	go sendValue(channel)
	go readValue(channel, done)
	<-done
}
