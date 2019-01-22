package main

import (
	"os"
	"log"
	"strings"
	"github.com/Shopify/sarama"
)

var (
	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	sarama.Logger = logger
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true                   //是否等待成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll          //等待服务器所有副本都保存成功后的响应
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机的分区类型

	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder("你好, 世界!")

	producer, err := sarama.NewSyncProducer(strings.Split("172.16.16.114:9092", ","), config)
	if err != nil {
		logger.Println("Failed to produce message,err = ", err)
		os.Exit(500)
	}
	defer producer.Close()

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Println("Failed to produce message,err = ", err)
	}
	logger.Printf("partition=%d, offset=%d\n", partition, offset)
}
