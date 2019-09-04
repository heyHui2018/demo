package main

import (
	"fmt"
	dis "github.com/heyHui2018/demo/etcd/discovery"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := dis.ServiceInfo{IP: "172.16.16.114"}

	s, err := dis.NewService(serviceName, serviceInfo, []string{
		"http://172.16.16.114:2379",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

	go func() {
		time.Sleep(time.Second * 20)
		s.Stop()
	}()

	s.Start()
}
