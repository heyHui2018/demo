package models

import (
	"github.com/heyHui2018/demo/gin/reserveProgram/base"
	"github.com/ngaut/log"
	"time"
)

type BindingLog struct {
	Id         int64
	OpenId     string    `xorm:"index 'openid'"    json:"openId"` // 服务号openid
	MoretvId   string    `xorm:"index 'moretv_id'" json:"moretvId"`
	CreateTime time.Time `xorm:"DateTime created 'create_time'"`
	BindStatus int       `xorm:"int(1)"            json:"bind_status"` // 服务号openid
}

func (this *BindingLog) InsertBindingLog() {
	_, err := base.DBEngine.Insert(this)
	if err != nil {
		log.Warnf("InsertBindingLog error, BindingLog = %+v,err = %v", this, err)
	} else {
		log.Infof("InsertBindingLog success, BindingLog = %+v", this)
	}
}
