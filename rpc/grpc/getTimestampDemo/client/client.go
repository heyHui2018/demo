package main

import (
	"context"
	"github.com/heyHui2018/best-practise/pb"
	"github.com/ngaut/log"
	"google.golang.org/grpc"
	"time"
)

const (
	address = "localhost:8667"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial error,err = %v", err)
	}
	defer conn.Close()
	c := pb.NewGetClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetTimestamp(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("error = %v", err)
	}
	log.Infof("Greply = %+v", r)
}
