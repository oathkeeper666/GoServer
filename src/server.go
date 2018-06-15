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
	var path = flag.String("c", "../etc/server.json", "server config path")
	var isDaemon = flag.Bool("d", false, "daemon flag")
	flag.Parse()

	if *isDaemon {
		if -1 == daemon() {
			fmt.Printf("create daemon process failed.\n")
			os.Exit(-1)
		}
	}

	network.StartListen()
	// load server config
	config.LoadServerConfig(*path)

	logger.WRITE_DEBUG("start server success!")
	waitForSignal()
}
