package main

import (
	"time"
	"encoding/json"
	"github.com/ngaut/log"
	"github.com/streadway/amqp"
)

var Channel *amqp.Channel
var Conn *amqp.Connection

func main() {
	exchangeName := "putong.exchange"
	routingKey := "putong.routingKey"
	MQStart(exchangeName, routingKey)

	msg := `{"msg":"hello world!"}`

	retryCount := 0
retry:
	if Channel == nil {
		Connect()
	}
	publish_info, err := json.Marshal(msg)
	if err != nil {
		log.Warnf("Publish Marshal msg error,err = %v,msg = %+v", err, msg)
		return
	}
	err = Channel.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: 2,
		Body:         publish_info,
	})
	if err != nil {
		log.Warn("Publish error,err = ", err)
		Channel.Close()
		Channel = nil
		Conn.Close()
		time.Sleep(3 * time.Second)
		if retryCount < 3 {
			retryCount++
			goto retry
		} else {
			log.Warnf("重试次数已满,不再重试")
		}
	} else {
		log.Info("Publish success")
	}
}

func MQStart(exchangeName, routingKey string) {
	Connect()
	deadExchangeName := "dead.exchange"
	deadQueueName := "dead.queue"
	deadRoutingKey := "dead.routingKey"
	queueName := "putong.queue"
	//死信
	err := Channel.ExchangeDeclare(deadExchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		log.Warnf("deadExchange 初始化 出错,err = %v", err)
	}
	_, err = Channel.QueueDeclare(deadQueueName, true, false, false, false, nil)
	if err != nil {
		log.Warnf("deadQueue 初始化 出错,err = %v", err)
	}
	err = Channel.QueueBind(deadQueueName, deadRoutingKey, deadExchangeName, false, nil)
	if err != nil {
		log.Warnf("绑定 %v 到 %v 出错,err = %v", deadQueueName, deadExchangeName, err)
	}
	//普通
	err = Channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		log.Warnf("exchange 初始化 出错,err = %v", err)
	}
	args := amqp.Table{}
	args["x-dead-letter-exchange"] = deadExchangeName
	args["x-dead-letter-routing-key"] = deadRoutingKey
	_, err = Channel.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		log.Warnf("queue 初始化 出错,err = %v", err)
	}
	err = Channel.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if err != nil {
		log.Warnf("绑定 %v 到 %v 出错,err = %v", queueName, exchangeName, err)
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
