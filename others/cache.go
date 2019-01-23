package main

import (
	"github.com/ngaut/log"
	"time"
	"sync"
)

var (
	_activityInfo string
	activityLock  sync.RWMutex
)

func GetActivityInfoCache() string {
	activityLock.RLock()
	defer activityLock.RUnlock()
	return _activityInfo
}

func Cache() {
	ActivityInfoCache()
	go func() {
		ticker1 := time.NewTicker(10 * time.Second)
		defer ticker1.Stop()
		for {
			select {
			case <-ticker1.C:
				{
					ActivityInfoCache()
				}
			}
		}
	}()
}

func ActivityInfoCache() {
	//do something
	reply := "111"
	activitySwap(reply)
}

func activitySwap(replyNew string) {
	activityLock.Lock()
	defer activityLock.Unlock()
	_activityInfo = replyNew
}

func main() {
	Cache()
	cache := GetActivityInfoCache()
	log.Infof("cache = %v", cache)
}
