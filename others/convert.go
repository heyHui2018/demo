package main

import (
	"github.com/ngaut/log"
	"strconv"
	"time"
)

func main() {
	//时间戳转指定格式字符串
	timestamp := int64(1548213455000)
	timeFormatStr := time.Unix(timestamp/1000, 0).Format("20060102")
	log.Infof("timeFormatStr = %v", timeFormatStr)

	//int64转string
	timeStr := strconv.FormatInt(timestamp, 10)
	log.Infof("timeStr = %v", timeStr)

	//string转int64
	timeInt64, err := strconv.ParseInt(timeStr+"999", 10, 64)
	if err != nil {
		log.Warnf("err = %v", err)
	}
	log.Infof("timeInt64 = %v", timeInt64)

	//string转float64
	var s string = "1.2453"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Warnf("err = %v", err)
	}
	log.Infof("f = %v", f)
}
