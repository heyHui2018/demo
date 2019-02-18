package main

import (
	"net"
	"time"
	"strings"
	"context"
	"encoding/json"
	"github.com/ngaut/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/heyHui2018/demo/rpc/grpc/getTimestampDemo"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetTimestamp(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	data := make(map[string]interface{})
	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	ymd := strings.Split(timeStr, "-")
	data["year"] = ymd[0]
	data["month"] = ymd[1]
	data["day"] = ymd[2]
	r := new(pb.GetReply)
	r.Status = 200
	r.Message = "success"
	r.Timestamp = time.Now().Unix()
	dataStr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	r.Data = string(dataStr)
	return r, nil
}

func main() {
	port := "8889"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("start listening on %v", port)
	s := grpc.NewServer()
	pb.RegisterGetServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
