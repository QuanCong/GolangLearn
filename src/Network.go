package main

import (
	"net"
	"log"
	"io"
	"time"
)
//tcp 服务端
func main() {

	listener, err := net.Listen("tcp", "localhost:9090")

	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleCoon(connection)

	}

}

func handleCoon(coon net.Conn){
	defer coon.Close()
	for{
		_,err:=io.WriteString(coon,time.Now().Format("15:04:05\n"))
		if err!=nil{
			return
		}
		time.Sleep(1*time.Second)
	}
}
