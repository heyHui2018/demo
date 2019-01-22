package base

import (
	"github.com/streadway/amqp"
	"github.com/ngaut/log"
	"time"
	"sync"
)

var MQWait *sync.WaitGroup
var BindConn *amqp.Connection
var BindChannel *amqp.Channel
var ProgramConn *amqp.Connection
var ProgramChannel *amqp.Channel

func BindConnect() {
	log.Info("BindConnect 开始连接")
	var err error
	username := GetConfig().MQs["bind"].Username
	password := GetConfig().MQs["bind"].Password
	ip := GetConfig().MQs["bind"].Ip
	port := GetConfig().MQs["bind"].Port
	host := GetConfig().MQs["bind"].Host
	mqUrl := "amqp://" + username + ":" + password + "@" + ip + ":" + port + "/" + host
a:
	BindConn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Warnf("BindConnect 连接MQ出错，err = %v", err)
		time.Sleep(3 * time.Second)
		goto a
	}

	BindChannel, err = BindConn.Channel()
	if err != nil {
		log.Warnf("BindConnect 打开channel出错，err = %v", err)
		BindConn.Close()
		time.Sleep(3 * time.Second)
		goto a
	}
	log.Info("BindConnect 连接完成")
}

func ProgramConnect() {
	log.Info("ProgramConnect 开始连接")
	var err error
	username := GetConfig().MQs["program"].Username
	password := GetConfig().MQs["program"].Password
	ip := GetConfig().MQs["program"].Ip
	port := GetConfig().MQs["program"].Port
	host := GetConfig().MQs["program"].Host
	mqUrl := "amqp://" + username + ":" + password + "@" + ip + ":" + port + "/" + host
a:
	ProgramConn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Warnf("ProgramConnect 连接MQ出错，err = %v", err)
		time.Sleep(3 * time.Second)
		goto a
	}

	ProgramChannel, err = ProgramConn.Channel()
	if err != nil {
		log.Warnf("ProgramConnect 打开channel出错，err = %v", err)
		ProgramConn.Close()
		time.Sleep(3 * time.Second)
		goto a
	}
	log.Info("ProgramConnect 连接完成")
}
