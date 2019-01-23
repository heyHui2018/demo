package main

import (
	"fmt"
	"time"
	"strings"
	"encoding/json"
	"github.com/ngaut/log"
	"github.com/garyburd/redigo/redis"
)

type TaskInfo struct {
	Sid     string `json:"sidinfo"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Date    string `json:"-"`
	Sort    int    `json:"sort"`
}

func main() {
	ip := "172.16.17.142"
	port := "6379"
	password := ""
	conn, err := GetRedisConnWithoutPool(ip, port, password)
	if err != nil {
		log.Warnf("connect to redis error,err = %v", err)
	}
	defer conn.Close()

	todayStr := time.Now().Format("2006-01-02")
	key := "2019CNYOperationActivity:" + todayStr
	sidListStr, err := redis.String(conn.Do("hget", key, "order"))
	if err != nil {
		log.Warnf("hget %v order error,err = %v", key, err)
		return
	}
	sidList := strings.Split(sidListStr, ",")

	var args []interface{}
	args = append(args, key) //放入key
	for k := range sidList {
		args = append(args, sidList[k]) //放入field
	}

	reply, err := redis.ByteSlices(conn.Do("hmget", args...))
	if err != nil {
		log.Warnf("hmget error,args = %v,err = %v", args, err)
		return
	}
	taskInfoList := make([]*models.TaskInfo, 0)
	for k := range reply {
		if nil == reply[k] {
			continue
		}
		taskInfo := new(TaskInfo)
		err = json.Unmarshal(reply[k], &taskInfo)
		if err != nil {
			log.Warnf("Unmarshal error,program = %+v,err = %v", reply[k], err)
			return
		}
		taskInfoList = append(taskInfoList, taskInfo)
	}
}

func GetRedisConnWithoutPool(ip, port, password string) (redis.Conn, error) {
	var conn redis.Conn
	var err error
	if len(password) > 0 {
		option := redis.DialPassword(password)
		conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", ip, port), option)
	} else {
		conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	}
	return conn, err
}
