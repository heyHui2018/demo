package main

import (
	"os"
	"net"
	"errors"
	"net/rpc"
	"syscall"
	"net/http"
	"os/signal"
	"github.com/ngaut/log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":12333")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	signs := make(chan os.Signal)
	signal.Notify(signs, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt, os.Kill, os.Interrupt)
	for {
		msg := <-signs
		log.Info("Recevied signal:", msg)
		clearWore()
	}
}

func clearWore() {
	//扫尾清理工作
	log.Info("开始停止程序....")
	log.Info("资源清理成功，退出")
	os.Exit(0)
}
