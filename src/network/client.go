package network

import (
	"net"
	"time"
	"sync"
	"logger"
	"servlet"
)

const (
	BUFSIZE = 1024
	MAX_CONN_COUNT = 8000
	TIME_OUT = 200				// 读写超时时间(ms)
)

/*
	保存session
*/
var sessionMap map[string]*Session = make(map[string]*Session)
var s_mutex *sync.Mutex = new(sync.Mutex)

func AddSession(s *Session) {
	s_mutex.Lock();
	defer s_mutex.Unlock()
	sessionMap[s.remoteIp] = s
}

func RemoveSession(ip string) {
	s_mutex.Lock();
	defer s_mutex.Unlock()
	delete(sessionMap, ip)
}

type Session struct {
	conn net.Conn
	remoteIp string
	handler servlet.Servlet
}

/*
	发送数据给客户端
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
	RemoveSession(this.remoteIp)
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
	go readCircle(s)

	return s
}

/*
	持续不断从客户端读取数据
*/
func readCircle(s *Session) {
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
			break
		}
	}
}

/*
	处理来自客户端的连接
*/
func HandleConnection(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(TIME_OUT * time.Microsecond));
	AddSession(NewSession(conn))
}