package main

import (
	"github.com/ngaut/log"
	"github.com/nsqio/go-nsq"
	"time"
)

func main() {
	producer, err := nsq.NewProducer("172.16.16.114:4150", nsq.NewConfig())
	if err != nil {
		log.Errorf("NewProducer error,err = %v", err)
	}
	if producer != nil {
		for i := 0; i < 5; i++ {
			err = producer.Publish("publish_test", []byte("publish_test"))
			if err != nil {
				log.Errorf("Publish error,err = %v", err)
			} else {
				log.Info("Publish successfully")
			}
			time.Sleep(3 * time.Second)
		}
	}
	producer.Stop()
}
