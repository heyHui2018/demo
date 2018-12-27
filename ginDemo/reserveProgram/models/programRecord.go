package models

import (
	"github.com/heyHui2018/demo/ginDemo/reserveProgram/base"
	"strconv"
	"time"
)

// UniqueKey: MoretvId+Sid
type ProgramRecord struct {
	Id          int64
	OpenId      string    `xorm:"index 'openid'"         json:"openId"` // 服务号openId
	MoretvId    string    `xorm:"index 'moretv_id'"      json:"moretvId"`
	Sid         string    `xorm:"index"                  json:"sid"`
	Pic         string    `xorm:""                       json:"pic"`
	Title       string    `xorm:""                       json:"title"`
	Type        string    `xorm:""                       json:"type"`    // 节目类型
	Remind      int       `xorm:"int(1)  index"          json:"remind"`  // 0 不 1 提醒
	Reserve     int       `xorm:"int(1)  index"          json:"reserve"` // 0 不 1 预约
	ReserveTime int64     `xorm:"bigint(20) index 'reserve_time'" json:"reserveTime"`
	CreateTime  time.Time `xorm:"DateTime created 'create_time'" json:"-"`
	UpdateTime  time.Time `xorm:"DateTime updated 'update_time'" json:"-"`
}

func (this *ProgramRecord) QueryByUniqueKey() (*ProgramRecord, error) {
	collectRecord := new(ProgramRecord)
	_, err := base.DBEngine.Where("moretv_id = ? and sid = ?", this.MoretvId, this.Sid).Get(collectRecord)
	if err != nil {
		return collectRecord, err
	}
	return collectRecord, nil
}

func (this *ProgramRecord) QueryByMoretvId() ([]*ProgramRecord, error) {
	collectRecordList := make([]*ProgramRecord, 0)
	sql := "select * from collect_record where moretv_id =" + this.MoretvId +
		" and sid in (select distinct(sid) from collect_record where moretv_id = " + this.MoretvId + " and (remind = 1 or reserve = 1)) order by update_time desc limit 100"
	results, err := base.DBEngine.Query(sql)
	if err != nil {
		return collectRecordList, err
	}
	for _, v := range results {
		collectRecord := new(ProgramRecord)
		collectRecord.OpenId = string(v["openid"])
		collectRecord.MoretvId = string(v["moretv_id"])
		collectRecord.Sid = string(v["sid"])
		collectRecord.Pic = string(v["pic"])
		collectRecord.Title = string(v["title"])
		collectRecord.Type = string(v["type"])
		remind, _ := strconv.Atoi(string(v["remind"]))
		collectRecord.Remind = remind
		reserve, _ := strconv.Atoi(string(v["reserve"]))
		collectRecord.Reserve = reserve
		reserveTime, _ := strconv.ParseInt(string(v["reserve_time"]), 10, 64)
		collectRecord.ReserveTime = reserveTime
		collectRecordList = append(collectRecordList, collectRecord)
	}
	return collectRecordList, nil
}

func (this *ProgramRecord) QueryByReserveAndReserveTime() ([]*ProgramRecord, error) {
	collectRecordList := make([]*ProgramRecord, 0)
	err := base.DBEngine.Where("reserve = ? and ? < reserve_time and reserve_time <= ?", this.Reserve, this.ReserveTime-60000, this.ReserveTime).Find(&collectRecordList)
	if err != nil {
		return collectRecordList, err
	}
	return collectRecordList, nil
}

func (this *ProgramRecord) QueryByProgramAndSid() ([]*ProgramRecord, error) {
	collectRecordList := make([]*ProgramRecord, 0)
	err := base.DBEngine.Where("remind = ? and sid = ?", this.Remind, this.Sid).Find(&collectRecordList)
	if err != nil {
		return collectRecordList, err
	}
	return collectRecordList, nil
}

func (this *ProgramRecord) UpdateRemindByUniqueKey() error {
	_, err := base.DBEngine.Cols("remind").Where("moretv_id = ? and sid = ?", this.MoretvId, this.Sid).Update(this)
	return err
}

func (this *ProgramRecord) UpdateReserveByUniqueKey() error {
	_, err := base.DBEngine.Cols("reserve", "reserve_time").Where("moretv_id = ? and sid = ?", this.MoretvId, this.Sid).Update(this)
	return err
}

func (this *ProgramRecord) UpdateReserveByPrimaryKey() error {
	_, err := base.DBEngine.Cols("reserve").Where("id = ?", this.Id).Update(this)
	return err
}

func (this *ProgramRecord) UpdateRemindByPrimaryKey() error {
	_, err := base.DBEngine.Cols("remind").Where("id = ?", this.Id).Update(this)
	return err
}

func (this *ProgramRecord) InsertCollectRecord() error {
	_, err := base.DBEngine.Insert(this)
	return err
}
