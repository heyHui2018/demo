package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"os"
	"sync"
)

var gLocker sync.Mutex; //全局锁
var gCondition *sync.Cond; //全局条件变量

var origin = "http://127.0.0.1:6666/"
var url = "ws://127.0.0.1:6666/echo"

//错误处理函数
func checkErr(err error, extra string) bool {
	if err != nil {
		formatStr := " Err : %s\n";
		if extra != "" {
			formatStr = extra + formatStr;
		}

		fmt.Fprintf(os.Stderr, formatStr, err.Error());
		return true;
	}

	return false;
}

//连接处理函数
func clientConnHandler(conn *websocket.Conn) {
	gLocker.Lock();
	defer gLocker.Unlock();
	defer conn.Close();
	request := make([]byte, 128);
	for {
		readLen, err := conn.Read(request)
		if checkErr(err, "Read") {
			gCondition.Signal();
			break;
		}

		//socket被关闭了
		if readLen == 0 {
			fmt.Println("Server connection close!");

			//条件变量同步通知
			gCondition.Signal();
			break;
		} else {
			//输出接收到的信息
			fmt.Println(string(request[:readLen]))

			//发送
			conn.Write([]byte("Hello !"));
		}

		request = make([]byte, 128);
	}
}

func main() {
	conn, err := websocket.Dial(url, "", origin);
	if checkErr(err, "Dial") {
		return;
	}

	gLocker.Lock();
	gCondition = sync.NewCond(&gLocker);
	_, err = conn.Write([]byte("Hello !"));
	go clientConnHandler(conn);

	//主线程阻塞，等待Singal结束
	for {
		//条件变量同步等待
		gCondition.Wait();
		break;
	}
	gLocker.Unlock();
	fmt.Println("Client finish.")
}

