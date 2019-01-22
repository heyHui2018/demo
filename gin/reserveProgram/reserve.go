package main

import (
	"github.com/heyHui2018/demo/gin/reserveProgram/service"
	"github.com/heyHui2018/demo/gin/reserveProgram/routers"
	"github.com/heyHui2018/demo/gin/reserveProgram/base"
	"github.com/ngaut/log"
	"os/signal"
	"net/http"
	"syscall"
	"fmt"
	"os"
)

func main() {
	base.ConfigInit()
	base.LogInit()
	base.DbInit()
	//go service.InformStart()
	service.MQStart()
	//service.TimedTask()
	//toolbox.StartTask()
	//go service.Monitor()
	//go beego.Run()

	routersInit := routers.InitRouter()
	httpPort := fmt.Sprintf(":%d", base.GetConfig().Server.HttpPort)
	//readTimeout := base.GetConfig().Server.ReadTimeout
	//writeTimeout := base.GetConfig().Server.WriteTimeout

	server := &http.Server{
		Addr:         httpPort,
		Handler:      routersInit,
		//ReadTimeout:  time.Duration(readTimeout),
		//WriteTimeout: time.Duration(writeTimeout),
	}
	log.Infof("start listening on %s", httpPort)
	go server.ListenAndServe()

	signs := make(chan os.Signal)
	signal.Notify(signs, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt, os.Kill, os.Interrupt)
	for {
		msg := <-signs
		log.Infof("Receive signal: %v", msg)
		clear()
	}
}

func clear() {
	log.Info("开始停止程序....")
	service.BindRegularCloseSign = true
	service.ProgramRegularCloseSign = true
	log.Info("资源清理成功，退出")
	os.Exit(0)
}
