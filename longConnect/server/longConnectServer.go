package demo

import (
	"net"
	"fmt"
	"os"
	"time"
)

func main() {
	server := ":7373"
	netListen, err := net.Listen("tcp", server)
	if err != nil {
		Log("connect error: ", err)
		os.Exit(1)
	}
	Log("Waiting for Client ...")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			Log(conn.RemoteAddr().String(), "Fatal error: ", err)
			continue
		}

		//设置短连接(10秒)
		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))

		Log(conn.RemoteAddr().String(), "connect success!")
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " Fatal error: ", err)
			return
		}

		Data := buffer[:n]
		message := make(chan byte)
		//心跳计时
		go HeartBeating(conn, message, 4)
		//检测每次是否有数据传入
		go GravelChannel(Data, message)

		Log(time.Now().Format("2006-01-02 15:04:05.0000000"), conn.RemoteAddr().String(), string(buffer[:n]))
	}
	defer conn.Close()
}

func GravelChannel(bytes []byte, mess chan byte) {
	for _, v := range bytes {
		mess <- v
	}
	close(mess)
}

func HeartBeating(conn net.Conn, bytes chan byte, timeout int) {
	select {
	case fk := <-bytes:
		Log(conn.RemoteAddr().String(), "心跳:第", string(fk), "times")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break

	case <-time.After(5 * time.Second):
		Log("conn dead now")
		conn.Close()
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
	return
}
