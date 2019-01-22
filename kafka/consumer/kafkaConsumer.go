package main

import (
	"os"
	"fmt"
	"log"
	"sync"
	"strings"
	"github.com/Shopify/sarama"
)

var (
	kafka  = "172.16.16.114:9092" //kafka服务器地址及端口号,此处可以指定多个地址,使用逗号分隔即可
	wg     sync.WaitGroup
	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	sarama.Logger = logger
	// 连接kafka
	consumer, err := sarama.NewConsumer(strings.Split(kafka, ","), nil)
	if err != nil {
		logger.Println("Failed to start consumer,err = ", err)
	}

	// consumer.Partitions 用户获取Topic上所有的Partitions。消息服务器上已经创建了test这个topic,因此此处指定参数为test.
	partitionList, err := consumer.Partitions("test")
	if err != nil {
		logger.Println("Failed to get the list of partitions,err = ", err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			logger.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Println("message is :", msg)
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}
	wg.Wait()
	logger.Println("Done consuming topic 'test'")
	consumer.Close()
}
