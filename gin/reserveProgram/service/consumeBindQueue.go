package service

import (
	"github.com/heyHui2018/demo/gin/reserveProgram/models"
	"github.com/heyHui2018/demo/gin/reserveProgram/utils"
	"github.com/heyHui2018/demo/gin/reserveProgram/base"
	"github.com/streadway/amqp"
	"github.com/ngaut/log"
	"encoding/json"
	"sync"
	"time"
)

var BindWait *sync.WaitGroup
var BindRegularCloseSign = false

func ConsumeBindQueue() {
	if base.BindChannel == nil {
		base.BindConnect()
	}
	err := base.BindChannel.Qos(1, 0, true)
	if err != nil {
		log.Warnf("ConsumeBindQueue 设置Qos出错，err = %v", err)
		base.BindChannel.Close()
		base.BindChannel = nil
		base.BindConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeBindQueue()
		return
	}
	msg, err := base.BindChannel.Consume(base.GetConfig().MQs["bind"].Queue, "", false, false, false, false, nil)
	if err != nil {
		log.Warn("ConsumeBindQueue 接收MQ消息出错，err = ", err)
		base.BindChannel.Close()
		base.BindChannel = nil
		base.BindConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeBindQueue()
		return
	}

	BindWait = new(sync.WaitGroup)
	for i := 0; i < base.GetConfig().MQs["bind"].ChanRangeNum; i++ {
		BindWait.Add(1)
		go rangeBindChannel(msg)
	}
	BindWait.Wait()

	if BindRegularCloseSign == true {
		log.Infof("BindRegularCloseSign = %v,ConsumeBindQueue已正常关闭", BindRegularCloseSign)
		base.MQWait.Done()
	} else {
		log.Infof("BindRegularCloseSign = %v,ConsumeBindQueue开始重连", BindRegularCloseSign)
		base.BindChannel.Close()
		base.BindChannel = nil
		base.BindConn.Close()
		go ConsumeBindQueue()
	}
}

func rangeBindChannel(msg <-chan amqp.Delivery) {
	defer BindWait.Done()
	for m := range msg {
		processId := time.Now().Format("20060102150405") + utils.GetRandString()
		binding := new(models.BindingMQData)
		err := json.Unmarshal(m.Body, binding)
		if err != nil {
			log.Warnf("rangeBindChannel Unmarshal msg 出错,err = %v,msgInfo = %v", err, string(m.Body))
		} else {
			log.Infof("rangeBindChannel processId = %v,binding = %+v", processId, binding)
			bindingRelation := new(models.BindingRelation)
			bindingRelation.MoretvId = binding.MoretvId
			bindingRelation.OpenId = binding.OpenId
			bindingLog := new(models.BindingLog)
			bindingLog.MoretvId = binding.MoretvId
			bindingLog.OpenId = binding.OpenId
			//查询moretvid是否已被绑定
			queryBR2, err := bindingRelation.QueryByMoretvId()
			if err != nil {
				log.Warnf("rangeBindChannel Query bindingRelation error,processId = %v,err = %v", processId, err)
			} else {
				if queryBR2.Id != 0 && queryBR2.OpenId != binding.OpenId {
					//moretvid已被绑定
					bindingLog.BindStatus = -1
					bindingLog.InsertBindingLog()
					bindingRelation.BindStatus = -1
					err := bindingRelation.UpdateBindStatusByUniqueKey()
					if err != nil {
						log.Warnf("rangeBindChannel UpdateBindStatusByUniqueKey error,processId = %v,err = %v", processId, err)
					}
				} else {
					//未被绑定
					bindingLog.BindStatus = 1
					bindingRelation.BindStatus = 1
					queryBR, err := bindingRelation.QueryByUniqueKey()
					if err != nil {
						log.Warnf("rangeBindChannel Query bindingRelation error,processId = %v,err = %v", processId, err)
					} else {
						if queryBR.Id != 0 {
							err = bindingRelation.UpdateMoretvIdByUniqueKey()
							if err != nil {
								log.Warnf("rangeBindChannel UpdateMoretvIdByUniqueKey error,processId = %v,err = %v", processId, err)
							} else {
								log.Infof("rangeBindChannel UpdateMoretvIdByUniqueKey success,processId = %v", processId)
								bindingLog.InsertBindingLog()
							}
						} else {
							bindingRelation.InsertBindingRelation()
							if err != nil {
								log.Warnf("rangeBindChannel InsertBindingRelation error, processId = %v,err = %v", processId, err)
							} else {
								log.Infof("rangeBindChannel InsertBindingRelation success, processId = %v", processId)
								bindingLog.InsertBindingLog()
							}
						}
					}
				}
			}
		}
		m.Ack(false)
	}
}
