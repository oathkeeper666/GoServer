package main

import (
	"fmt"
	"os"
	"network"
	"config"
	"logger"
	"flag"
	"mysqldb"
	"util"
) 

func main() {
	// parse cmd flag
	var path = flag.String("c", "../etc/game.json", "server config path")
	var isDaemon = flag.Bool("d", false, "daemon flag")
	flag.Parse()

	// daemon process
	if *isDaemon {
		if -1 == util.Daemon() {
			fmt.Printf("create daemon process failed.\n")
			os.Exit(-1)
		}
	}

	// load server config
	config.LoadGameConfig(*path)

	// new logger
	logger.NewLog(config.GameConf.LogPath, uint8(config.GameConf.MinLevel), *isDaemon)

	// listen
	if !network.StartListen(config.GameConf.Protocol, config.GameConf.ListenAddress) {
		logger.WRITE_ERROR("listen for client failed.")
		return
	}

	// handler msg
	// todo

	// open database
	if mysqldb.ConnectToDb() {
		logger.WRITE_DEBUG("open database success.")
	}

	logger.WRITE_DEBUG("start server success!")

	// wait to exit
	util.WaitForSignal()

	// main goroutine run
	network.H.Run()
}
