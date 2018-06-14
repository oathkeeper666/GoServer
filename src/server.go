package main

import (
	"fmt"
	"syscall"
	"os"
	"os/signal"
	"network"
	"common"
	"config"
	"logger"
) 

func waitForSignal() {
	signals := [] os.Signal {
		syscall.SIGINT,
		syscall.SIGQUIT,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals ...)
	s := <-c
	fmt.Printf("process receive signal %v\n", s)
}

func main() {
	network.StartListen()
	fmt.Println(common.LOGIN_RESPOND)
	// load server config
	config.LoadServerConfig("../etc/server.json")
	log := logger.GetLog("server")
	if log != nil {
		log.WriteLog(logger.INFO, "Number one is %d", 1)
	}

	waitForSignal()
}
