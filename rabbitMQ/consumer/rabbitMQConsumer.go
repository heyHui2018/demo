package main

import (
	"github.com/ngaut/log"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var Channel *amqp.Channel
var Conn *amqp.Connection
var Sign = false
var RangeWait *sync.WaitGroup
var CloseWait *sync.WaitGroup

func Consumer() {
	queue := "putong.queue"

	MQStart(queue)

	rangeNum := 1
	if Channel == nil {
		Connect()
	}
	err := Channel.Qos(1, 0, true)
	if err != nil {
		log.Warnf("Consumer 设置Qos出错，err = %v", err)
		Channel.Close()
		Channel = nil
		Conn.Close()
		CloseWait.Done()
		time.Sleep(3 * time.Second)
		go Consumer()
		return
	}
	msg, err := Channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		log.Warn("Consumer 接收MQ消息出错，err = ", err)
		Channel.Close()
		Channel = nil
		Conn.Close()
		CloseWait.Done()
		time.Sleep(3 * time.Second)
		// 因queue被删除会导致在次循环出错,故此出错重试逻辑需能重建exchange及queue及bind,此例中Consumer()会调用MQStart(queue),满足此要求
		go Consumer()
		return
	}

	RangeWait = new(sync.WaitGroup)
	for i := 0; i < rangeNum; i++ {
		RangeWait.Add(1)
		go rangeChannel(msg)
	}

	RangeWait.Wait()

	if Sign == true {
		log.Infof("Sign = %v,Consumer已正常关闭", Sign)
		CloseWait.Done()
	} else {
		log.Infof("Sign = %v,Consumer开始重连", Sign)
		Channel.Close()
		Channel = nil
		Conn.Close()
		go Consumer()
	}
}

func rangeChannel(msg <-chan amqp.Delivery) {
	defer RangeWait.Done()
	for m := range msg {
		// do something
		log.Infof("rangeChannel m = %+v, body = %v", m, string(m.Body))
		// m.Reject(false)
		// m.Ack(false)
		m.Reject(false)
	}
}

func MQStart(queueName string) {
	Connect()

	Exchange := "putong.exchange"
	err := Channel.ExchangeDeclare(Exchange, "direct", true, false, false, false, nil)
	if nil != err {
		log.Warnf("MQStart 初始化 Exchange:%v 出错,err = %v", Exchange, err)
	}

	// _, err = Channel.QueueDeclare(queueName, true, false, false, false, nil)
	// if err != nil {
	// 	log.Warnf("MQStart 初始化 Queue:%v 出错,err = %v", queueName, err)
	// }

	routingKey := "putong.routingKey"
	err = Channel.QueueBind(queueName, routingKey, Exchange, false, nil)
	if err != nil {
		log.Warnf("MQStart 绑定 %v 到 %v 出错,err = %v", queueName, Exchange, err)
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
		log.Warnf("Connect 连接MQ出错,err = %v", err)
		time.Sleep(1 * time.Second)
		goto a
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Warnf("Connect 打开channel出错,err = %v", err)
		Conn.Close()
		time.Sleep(1 * time.Second)
		goto a
	}
	log.Info("Connect 连接完成")
}

func main() {
	CloseWait = new(sync.WaitGroup)
	CloseWait.Add(1)
	Consumer()
	CloseWait.Wait()
}
