package network

import (
	"net"
	"logger"
	"servlet"
)

const (
	BUFSIZE = 1024
)

type Session struct {
	conn net.Conn
	remoteIp string
	handler servlet.Servlet
}

/*
	发送数据给对端
*/
func (this *Session) SendData(data []byte) {
	_, err := this.conn.Write(data)
	if err != nil {
		logger.WRITE_WARNING("write to %s error: %v", this.remoteIp, err)
		this.Close()
	}
}

/*
	处理从客户端那里读取过来的数据
*/
func (this *Session) ProcessData(data []byte) {
	
}

/*
	设置处理协议的处理对象
*/
func (this *Session) SetHandler() {
	// this.handler = new(servlet.LoginSevlet) 
	this.handler = servlet.NewLoginServlet()

}

/*
	关闭与客户端通信的session，执行清理操作
*/
func (this *Session) Close() {
	this.conn.Close()
}

/*
	持续不断从客户端读取数据
*/
func (s *Session) readCircle() {
	for {
		b := make([]byte, BUFSIZE)
		_, err := s.conn.Read(b)
		if err == nil {
			// do something
			logger.WRITE_DEBUG("read data from %s: %s", s.remoteIp, b)
			s.ProcessData(b)
		} else {
			logger.WRITE_WARNING("read from %s error: %v", s.remoteIp, err)
			s.Close()
			RemoveSession(s.remoteIp)
			break
		}
	}
}

/*
	创建一个与客户端通信的session
*/
func NewSession(conn net.Conn) (*Session) {
	s := &Session {
		conn: conn,
		remoteIp: conn.RemoteAddr().String(),
	}
	s.SetHandler();
	go s.readCircle()

	return s
}