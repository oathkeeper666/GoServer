package network

import (
	"net"
	"time"
	//"strings"
	"config"
	"logger"
	"servlet"
)

const (
	TIME_OUT = 100				// 读写超时时间(ms)
)

var listener net.Listener
var ConnCh = make(chan net.Conn, 10)

func GetHostAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.WRITE_ERROR("get local host address error, err is %v\n.", err)
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalUnicast() {
			return ipnet.IP.String()
		}
	}

	return ""
}

func StartListen(protocol string, address string) bool {
	/*ip := GetHostAddr()
	port := "9000"

	var err error
	listener, err = net.Listen("tcp", strings.Join([]string { ip, port }, ":"))
	if err != nil {
		logger.WRITE_ERROR("Listening error, error is %v\n", err)
		return false;
	}*/

	var err error	
	listener, err = net.Listen(protocol, address)
	if err != nil {
		logger.WRITE_ERROR("Listening error, error is %v\n", err)
		return false;
	}
	
	// start accept
	go func() {
		logger.WRITE_DEBUG("begin to accepting ...") 
		for {
			conn, err := listener.Accept()
			if err != nil {
				logger.WRITE_WARNING("accept error, error is %v\n", err)
			}
			ConnCh <- conn
		}
	}()

	return true
}