package service

import (
	"github.com/heyHui2018/demo/ginDemo/reserveProgram/models"
	"github.com/heyHui2018/demo/ginDemo/reserveProgram/utils"
	"github.com/heyHui2018/demo/ginDemo/reserveProgram/base"
	"github.com/streadway/amqp"
	"github.com/ngaut/log"
	"encoding/json"
	"sync"
	"time"
)

var ProgramWait *sync.WaitGroup
var ProgramRegularCloseSign = false

func ConsumeProgramQueue() {
	if base.ProgramChannel == nil {
		base.ProgramConnect()
	}
	err := base.ProgramChannel.Qos(1, 0, true)
	if err != nil {
		log.Warnf("ConsumeProgramQueue 设置Qos出错，err = %v", err)
		base.ProgramChannel.Close()
		base.ProgramChannel = nil
		base.ProgramConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeProgramQueue()
		return
	}
	msg, err := base.ProgramChannel.Consume(base.GetConfig().MQs["program"].Queue, "", false, false, false, false, nil)
	if err != nil {
		log.Warn("ConsumeProgramQueue 接收MQ消息出错，err = ", err)
		base.ProgramChannel.Close()
		base.ProgramChannel = nil
		base.ProgramConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeProgramQueue()
		return
	}

	ProgramWait = new(sync.WaitGroup)
	for i := 0; i < base.GetConfig().MQs["program"].ChanRangeNum; i++ {
		ProgramWait.Add(1)
		go rangeProgramChannel(msg)
	}
	ProgramWait.Wait()

	if ProgramRegularCloseSign == true {
		log.Infof("ProgramRegularCloseSign = %v,ConsumeProgramQueue已正常关闭", ProgramRegularCloseSign)
		base.MQWait.Done()
	} else {
		log.Infof("ProgramRegularCloseSign = %v,ConsumeProgramQueue开始重连", ProgramRegularCloseSign)
		base.ProgramChannel.Close()
		base.ProgramChannel = nil
		base.ProgramConn.Close()
		go ConsumeProgramQueue()
	}
}

func rangeProgramChannel(msg <-chan amqp.Delivery) {
	defer ProgramWait.Done()
	for m := range msg {
		start := time.Now()
		processId := time.Now().Format("20060102150405") + utils.GetRandString()
		program := new(models.ProgramMQData)
		err := json.Unmarshal(m.Body, program)
		if err != nil {
			log.Warnf("rangeProgramChannel Unmarshal msg 出错,err = %v,msgInfo = %v", err, string(m.Body))
		} else {
			log.Infof("rangeProgramChannel processId = %v,program = %+v", processId, program)
			programRecord := new(models.ProgramRecord)
			programRecord.Remind = 1
			programRecord.Sid = program.Sid
			var programRecordList []*models.ProgramRecord
			for i := 0; i < 3; i++ {
				programRecordList, err = programRecord.QueryByProgramAndSid()
				if err != nil {
					log.Warnf("rangeRemindChannel QueryByRemindAndSid error,processId = %v, err = %v", processId, err)
					time.Sleep(5 * time.Second)
					continue
				}
				break
			}
			if len(programRecordList) == 0 {
				log.Infof("rangeRemindChannel 此节目无人设置提醒,sid = %v,title = %v", programRecord.Sid, program.Title)
			} else {
				count := 0
				//for _, v := range programRecordList {
				//	count++
					//ProgramChan <- v
				//}
				log.Infof("rangeRemindChannel 提醒通知执行完成,执行数量 = %v,执行耗时 = %v,sid = %v", count, time.Since(start), programRecord.Sid)
			}
		}
		m.Ack(false)
	}
}
