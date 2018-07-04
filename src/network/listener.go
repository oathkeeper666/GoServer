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
	/*ip := GetHostAddr()
	port := "9000"

	var err error
	listener, err = net.Listen("tcp", strings.Join([]string { ip, port }, ":"))
	if err != nil {
		logger.WRITE_ERROR("Listening error, error is %v\n", err)
		return false;
	}*/

	var err error	
	listener, err = net.Listen("tcp", config.SrvConf.ListenAddress)
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
			HandleConnection(conn, getServerServlet())
		}
	}()

	return true
}

/*
	建立客户端协议处理链
*/
func buildClientServlet() servlet.Servlet {
	login_servlet := servlet.NewLoginServlet()
	return login_servlet
}

/*
	建立内部服务器协议处理链
*/
func buildPeerServlet() servlet.Servlet {
	return nil
}

func getServerServlet() servlet.Servlet {
	if config.SrvConf.ServerId == "wb-game-server" {
		return buildClientServlet()
	} else {
		return buildPeerServlet()
	}
}

/*
	处理来自客户端的连接
*/
func HandleConnection(conn net.Conn, servlet servlet.Servlet) {
	conn.SetDeadline(time.Now().Add(TIME_OUT * time.Microsecond));
	s := NewSession(conn)
	s.SetHandler(servlet)
	// save session
	AddSession(s)
}