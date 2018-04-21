package network

import (
	"net"
	"fmt"
	"strings"
)

var listener net.Listener

func GetHostAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("get local host address error, err is %v\n.", err)
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalUnicast() {
			return ipnet.IP.String()
		}
	}

	return ""
}

func StartListen() {
	ip := GetHostAddr()
	port := "9000"

	var err error
	listener, err = net.Listen("tcp", strings.Join([]string { ip, port }, ":"))
	if err != nil {
		fmt.Printf("Listening error, error is %v\n", err)
		return
	}
	
	// start accept
	go func() {
		fmt.Println("begin to accepting ...") 
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Printf("accept error, error is %v\n", err)
			}
			HandleConnection(NewClient(conn))
		}
	}()
}