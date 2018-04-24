package main

import (
	"fmt"
	"syscall"
	"os"
	"os/signal"
	"network"
	"common"
	"config"
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

var str string = `{
	"ServerId": "wb-game-server1",
	"LogPath": "log/"
}`

func main() {
	network.StartListen()
	fmt.Println(common.LOGIN_RESPOND)

	srvConf := config.FromXmlFile("../etc/tasks.xml")
	if srvConf != nil {
		fmt.Printf("server xml is %v.\n", srvConf)
	}

	waitForSignal()
}
