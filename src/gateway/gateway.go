package gateway

import (
	"fmt"
	"os"
	"network"
	"config"
	"logger"
	"flag"
	"mysqldb"
	"util"
	"gateway/servlet"
) 

func buildMsgHandler() {
	login_servlet := servlet.NewLoginServlet()
	network.SetHandler(login_servlet)
}

func main() {
	// parse cmd flag
	var path = flag.String("c", "../etc/gateway.json", "server config path")
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
	config.LoadGatewayConfig(*path)

	// new logger
	logger.NewLog(config.GatewayConf.LogPath, config.GatewayConf.MinLevel, *isDaemon)

	// listen
	if !network.StartListen(config.GatewayConf.Protocol, config.GatewayConf.ListenAddress) {
		logger.WRITE_ERROR("listen for client failed.");
		return
	}
	buildMsgHandler()
	
	// open database
	if mysqldb.ConnectToDb() {
		logger.WRITE_DEBUG("open database success.")
	}

	// main goroutine run
	network.H.run()

	// wait to exit
	util.WaitForSignal()

	logger.WRITE_DEBUG("start server success!")
}
