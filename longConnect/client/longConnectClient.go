package main

import (
	"net"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	server := "127.0.0.1:7373"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		Log(os.Stderr, "Fatal error:", err.Error())
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		Log("Fatal error:", err.Error())
		return
	}
	Log(conn.RemoteAddr().String(), "connection succcess!")

	sender(conn)
	Log("send over")
}

func sender(conn *net.TCPConn) {
	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " Hello I'm Heartbeat Client."
		msg, err := conn.Write([]byte(words))
		if err != nil {
			Log(conn.RemoteAddr().String(), "Fatal error: ", err)
			os.Exit(1)
		}
		Log("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}
	for i := 0; i < 2; i++ {
		time.Sleep(12 * time.Second)
	}
	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " Hi I'm Heartbeat Client."
		msg, err := conn.Write([]byte(words))
		if err != nil {
			Log(conn.RemoteAddr().String(), "Fatal error: ", err)
			os.Exit(1)
		}
		Log("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
	return
}
