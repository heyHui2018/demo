package demo

import (
	"fmt"
	"time"
)

//ch堵塞，当timeout有数据时，执行case <-timeout:分支，实现超时控制
func mian() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")
	}
}
