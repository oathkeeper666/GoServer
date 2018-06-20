package network

import (
	"net"
	"strings"
	// "config"
	"logger"
)

var listener net.Listener

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

func StartListen() bool {
	ip := GetHostAddr()
	port := "9000"

	var err error
	listener, err = net.Listen("tcp", strings.Join([]string { ip, port }, ":"))
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
			HandleConnection(conn)
		}
	}()

	return true
}