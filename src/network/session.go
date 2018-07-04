package network

import (
	"net"
	"logger"
	"servlet"
	"encoding/binary"
	"bytes"
	"sync"
)

const (
	BUFSIZE = 1024
)

/*
	服务器与对端进行通信的结构
*/
type Session struct {
	conn net.Conn
	remoteIp string
	sid int64
	pid int64
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
func (this *Session) ProcessData(num int32, data []byte) {
	if this.handler != nil {
		this.handler.HandleMsg(this.sid, num, data)
	}
}

/*
	设置处理协议的处理对象
*/
func (this *Session) SetHandler(servlet servlet.Servlet) {
	this.handler = servlet
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
		b := make([]byte, BUFSIZE)	// 网络字节序
		_, err := s.conn.Read(b)
		if err == nil {
			logger.WRITE_DEBUG("read data from %s", s.remoteIp)
			// 获取协议号
			num, num_err := convertNum(b[:4])
			if num_err != nil {
				logger.WRITE_ERROR("convert proto num error: %v", num_err)
				continue
			}
			// 网络字节序转化为本机字节序
			c := make([]byte, BUFSIZE)
			conv_err := convertToHost(b, c)
			if conv_err != nil {
				logger.WRITE_ERROR("convert network byte stream to local byte error: %v", conv_err)
				continue
			}
			// 处理协议
			s.ProcessData(num, c)
		} else {
			logger.WRITE_WARNING("read from %s error: %v", s.remoteIp, err)
			s.Close()
			RemoveSession(s)
			break
		}
	}
}

/*
	将byte转化为int
*/
func convertNum(b []byte) (int32, error) {
	var num int32
	reader := bytes.NewReader(b)	
	err := binary.Read(reader, binary.BigEndian, &num)
	return num, err
}

/*
	网络字节序转化为本机字节序
*/
func convertToHost(b []byte, c []byte) error {
	reader := bytes.NewReader(b)
	return binary.Read(reader, binary.BigEndian, c)
}


/*
	获得一个sid
*/
var g_sid int64 = 0
var sid_mutex *sync.Mutex = new(sync.Mutex)
func getSid() int64 {
	sid_mutex.Lock()
	defer sid_mutex.Unlock()
	g_sid = g_sid + 1
	return g_sid
}

/*
	创建一个与客户端通信的session
*/

func NewSession(conn net.Conn) (*Session) {
	s := &Session {
		conn: conn,
		remoteIp: conn.RemoteAddr().String(),
		sid: getSid(),
		pid: 0,
	}
	go s.readCircle()

	return s
}