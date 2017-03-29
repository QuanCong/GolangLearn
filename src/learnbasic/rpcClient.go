package main

import (
	"net/rpc"
	"fmt"
	"rpcService"
)

func main() {
	var severAddress = "localhost"
	client, err := rpc.DialHTTP("tcp", severAddress + ":1234")
	if err != nil {
		fmt.Println(err)
	}

	args := &rpcService.Args{10, 12}
	var reply int
	/*client.Call("Calculate.Add", args, &reply)
	fmt.Println("result:", reply)
*/

	//异步方式
	replyCall:=client.Go("Calculate.Add",args,&reply, nil)
	fmt.Println("调用返回...")
	<-replyCall.Done
	fmt.Println("result:",reply)



}
