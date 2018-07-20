package main

import (
	"fmt"
	"syscall"
	"os"
	"os/signal"
	"network"
	"config"
	"logger"
	"flag"
	"mysqldb"
	"util"
) 

func waitForSignal() {
	signals := [] os.Signal {
		syscall.SIGINT,
		syscall.SIGQUIT,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals ...)
	s := <-c
	fmt.Printf("process exit, receive signal %v\n", s)
	onExit()
}

// 进程退出时执行的清理操作
func onExit() {
	mysqldb.CloseDB()
}

func daemon() int {
	// fork
	ret, ret2, err := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if ret2 < 0 || err != 0 {
		fmt.Printf("fork failed: %v, %v\n", ret2, err)
		return -1
	}

	// parent exit
	if ret > 0 {
		os.Exit(0)
	}

	// change file mode mask
	syscall.Umask(0)

	// set sid
	_, serr := syscall.Setsid()
	if serr != nil {
		fmt.Printf("set sid failed, err is %v\n", serr)
		return -1
	}

	// close stdin, stdout, stderr
	f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if e == nil {
		fd := f.Fd()
		syscall.Dup2(int(fd), int(os.Stdin.Fd()))
		syscall.Dup2(int(fd), int(os.Stdout.Fd()))
		syscall.Dup2(int(fd), int(os.Stderr.Fd()))
	}

	return 0
}

func main() {
	// parse cmd flag
	var path = flag.String("c", "../etc/gateway.json", "server config path")
	var isDaemon = flag.Bool("d", false, "daemon flag")
	flag.Parse()

	// daemon process
	if *isDaemon {
		if -1 == daemon() {
			fmt.Printf("create daemon process failed.\n")
			os.Exit(-1)
		}
	}

	// load server config
	config.LoadGatewayConfig(*path)

	// new logger
	logger.NewLog(config.GatewayConf.LogPath, logger.DEBUG, *isDaemon)

	// listen
	if !network.StartListen() {
		logger.WRITE_ERROR("listen for client failed.");
		return
	}
	logger.WRITE_DEBUG("start server success!")

	// open database
	if mysqldb.ConnectToDb() {
		logger.WRITE_DEBUG("open database success.")
	}

	// wait to exit
	waitForSignal()
}
