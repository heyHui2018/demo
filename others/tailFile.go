package main

import (
	"github.com/hpcloud/tail"
	"github.com/ngaut/log"
)

func main() {
	offSet := 0
	seekInfo := new(tail.SeekInfo)
	seekInfo.Offset = offSet
	seekInfo.Whence = 0//0-从头 1-从当前位置 2-从尾
	log.Infof("seekInfo = %+v", seekInfo)
	t, err := tail.TailFile("test3.log", tail.Config{Follow: true, Location: seekInfo})
	if err != nil {
		log.Warnf("err = %v", err)
		return
	}
	for line := range t.Lines {
		curOffSet, err := t.Tell()
		if err != nil {
			log.Warnf("111err = %v", err)
		}
		log.Infof("test = %v,offSet = %v", line.Text, curOffSet)
		// 可在此添加偏移量存储逻辑,以防程序退出
	}
}
