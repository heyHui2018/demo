package service

import (
	"github.com/heyHui2018/demo/ginDemo/reserveProgram/base"
	"github.com/ngaut/log"
	"sync"
)

func MQStart() {
	base.BindConnect()
	base.ProgramConnect()
	bindExchange := base.GetConfig().MQs["bind"].Exchange
	err := base.BindChannel.ExchangeDeclare(bindExchange, "direct", true, false, false, false, nil)
	if nil != err {
		log.Warnf("MQInit 初始化 Exchange:%v 出错,err = %v", bindExchange, err)
	}
	programExchange := base.GetConfig().MQs["program"].Exchange
	err = base.ProgramChannel.ExchangeDeclare(programExchange, "topic", true, false, false, false, nil)
	if nil != err {
		log.Warnf("MQInit 初始化 Exchange:%v 出错,err = %v", programExchange, err)
	}

	bindQueue := base.GetConfig().MQs["bind"].Queue
	_, err = base.BindChannel.QueueDeclare(bindQueue, true, false, false, false, nil)
	if err != nil {
		log.Warnf("MQInit 初始化 Queue:%v 出错,err = %v", bindQueue, err)
	}
	programQueue := base.GetConfig().MQs["program"].Queue
	_, err = base.ProgramChannel.QueueDeclare(programQueue, true, false, false, false, nil)
	if err != nil {
		log.Warnf("MQInit 初始化 Queue:%v 出错,err = %v", programQueue, err)
	}

	bindKey := base.GetConfig().MQs["bind"].Key
	err = base.BindChannel.QueueBind(bindQueue, bindKey, bindExchange, false, nil)
	if err != nil {
		log.Warnf("MQInit 绑定 %v 到 %v 出错,err = %v", bindQueue, bindExchange, err)
	}
	programKey := base.GetConfig().MQs["program"].Key
	err = base.ProgramChannel.QueueBind(programQueue, programKey, programExchange, false, nil)
	if err != nil {
		log.Warnf("MQInit 绑定 %v 到 %v 出错,err = %v", programQueue, programExchange, err)
	}
	log.Info("MQInit完成")
	base.MQWait = new(sync.WaitGroup)
	base.MQWait.Add(2)
	go ConsumeBindQueue()
	go ConsumeProgramQueue()
}