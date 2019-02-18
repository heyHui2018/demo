package main

import (
	"net/rpc"
	"github.com/ngaut/log"
)

type Args struct {
	A, B int
}

func main() {
	serverAddress := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", serverAddress+":12333")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("call error:", err)
	}
	log.Infof("Arith: %d*%d=%d", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	log.Infof("replyCall = %v", replyCall)
}
