package main

import (
	"time"
	"encoding/json"
	"github.com/ngaut/log"
	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func main() {
	MQStart()

	msg := `{"msg":"hello world!"}`
	exchange := "direct_exchange"
	routingKey := "test_routing_key"

	retryCount := 0
retry:
	if Channel == nil {
		Connect()
	}
	publishInfo, err := json.Marshal(msg)
	if err != nil {
		log.Warnf("main Marshal error,err = %v", err)
		return
	}
	err = Channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		//DeliveryMode: 2,
		//exchange/queue均为durable时，在此设置DeliveryMode为2即可持久化数据
		Body: publishInfo,
	})
	if err != nil {
		log.Warnf("Publish error,err = %v", err)
		Channel.Close()
		Channel = nil
		Conn.Close()
		time.Sleep(3 * time.Second)
		if retryCount < 3 {
			retryCount++
			goto retry
		} else {
			log.Warnf("重试次数已满，不再重试")
		}
	} else {
		log.Info("Publish success")
	}
}

func MQStart() {
	Connect()

	Exchange := "direct_exchange"
	err := Channel.ExchangeDeclare(Exchange, "direct", true, false, false, false, nil)
	if nil != err {
		log.Warnf("MQStart 初始化 Exchange:%v 出错,err = %v", Exchange, err)
	}

	log.Info("MQStart完成")
}

func Connect() {
	log.Info("Connect 开始连接")
	var err error
	username := "guest"
	password := "guest"
	ip := "127.0.0.1"
	port := "5672"
	host := "test_host"
	mqUrl := "amqp://" + username + ":" + password + "@" + ip + ":" + port + "/" + host
a:
	Conn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Warnf("Connect 连接MQ出错，err = %v", err)
		time.Sleep(3 * time.Second)
		goto a
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Warnf("Connect 打开channel出错，err = %v", err)
		Conn.Close()
		time.Sleep(3 * time.Second)
		goto a
	}
	log.Info("Connect 连接完成")
}
