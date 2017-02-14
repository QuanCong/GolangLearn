package main

import (
	"rpcService"
	"net/rpc"
	"net"
	"fmt"
	"net/http"
	"time"
)

func main() {
	caculate := new(rpcService.Calculate)
	rpc.Register(caculate)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		fmt.Println(e)
	}
	go http.Serve(l, nil)
	fmt.Println("rpc server started")
	for ; ; {
		time.Sleep(1000 * time.Second)
	}

}
