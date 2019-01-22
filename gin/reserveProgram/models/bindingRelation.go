package models

import (
	"github.com/heyHui2018/demo/gin/reserveProgram/base"
	"time"
)

type BindingRelation struct {
	Id         int64
	OpenId     string    `xorm:"unique 'openid'"    json:"openId"` // 服务号openid
	MoretvId   string    `xorm:"index  'moretv_id'" json:"moretvId"`
	BindStatus int       `xorm:"int(1)"             json:"bind_status"` // 服务号openid
	CreateTime time.Time `xorm:"DateTime created 'create_time'"`
	UpdateTime time.Time `xorm:"DateTime updated 'update_time'"`
}

func (this *BindingRelation) QueryByUniqueKey() (*BindingRelation, error) {
	bindingRelation := new(BindingRelation)
	_, err := base.DBEngine.Where("openid = ?", this.OpenId).Get(bindingRelation)
	if err != nil {
		return bindingRelation, err
	}
	return bindingRelation, nil
}

func (this *BindingRelation) QueryByMoretvId() (*BindingRelation, error) {
	bindingRelation := new(BindingRelation)
	_, err := base.DBEngine.Where("moretv_id = ?", this.MoretvId).Get(bindingRelation)
	if err != nil {
		return bindingRelation, err
	}
	return bindingRelation, nil
}

func (this *BindingRelation) UpdateMoretvIdByUniqueKey() error {
	_, err := base.DBEngine.Cols("moretv_id", "bind_status").Where("openid = ?", this.OpenId).Update(this)
	return err
}

func (this *BindingRelation) UpdateBindStatusByUniqueKey() error {
	_, err := base.DBEngine.Cols("bind_status").Where("openid = ?", this.OpenId).Update(this)
	return err
}

func (this *BindingRelation) InsertBindingRelation() error {
	_, err := base.DBEngine.Insert(this)
	return err
}
