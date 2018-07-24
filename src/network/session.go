package network

import (
	"net"
	"logger"
	"servlet"
	"encoding/binary"
	"bytes"
	"consts"
	"time"
	"fmt"
	"util"
)

type Package struct {
	sid string
	cmd int32
	data []byte
}

func NewPackage() *Package {
	return &Package {
		sid: "",
		cmd: 0,
		data: nil,
	}
}

/*
	服务器与对端进行通信的结构
*/
type Session struct {
	conn net.Conn
	remoteIp string
	sid string
	lastHeartBeateTime time.Time
	heartBeateDuration time.Duration
	hub *Hub 	
}

// session处理协议的Handler
var handler servlet.Servlet

/*
	处理从客户端那里读取过来的数据
*/
func (this *Session) ProcessData(pack *Package) {
	this.lastHeartBeateTime = time.Now()
	if handler != nil {
		op := util.StartOperation("ProcessData")
		ret_msg := handler.HandleMsg(pack.cmd, pack.data)
		this.SendToPeer(ret_msg)
		op.Finish(time.Millisecond * 100)
	}
}

/*
	给对端发送数据
*/
func (this *Session) SendToPeer(msg []byte) {
	if len(msg) > consts.SESSION_WRITE_MSG_SIZE {
		logger.WRITE_WARNING("send data to peer failed, msg size too big.")
		return
	}
	_, err := this.conn.Write(msg)
	if err != nil {
		logger.WRITE_WARNING("write data to peer error: %v", err)
		this.hub.unregister <- this
	}
}

/*
	设置处理协议的处理对象
*/
func SetHandler(servlet servlet.Servlet) {
	handler = servlet
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
		b := make([]byte, consts.SESSION_READ_MSG_SIZE)	// 网络字节序
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
			c := make([]byte, consts.SESSION_READ_MSG_SIZE)
			conv_err := convertToHost(b, c)
			if conv_err != nil {
				logger.WRITE_ERROR("convert network byte stream to local byte error: %v", conv_err)
				continue
			}

			// 缓存协议
			pack := NewPackage()
			pack.sid = s.sid
			pack.cmd = num
			pack.data = c
			s.hub.receiveMsgQueue <- pack
		} else {
			logger.WRITE_WARNING("read from %s error: %v", s.remoteIp, err)
			s.hub.unregister <- s
			break
		}
	}
}

func (s *Session) String() string {
	return fmt.Sprintf("%d@%s", s.sid, s.conn.RemoteAddr())
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
	创建一个与客户端通信的session
*/
func NewSession(conn net.Conn) (*Session) {
	s := &Session {
		conn: conn,
		remoteIp: conn.RemoteAddr().String(),
		sid: util.GenUUID(),
		lastHeartBeateTime: time.Now(),
		heartBeateDuration: consts.SESSION_HEART_BEATE_INTERVAL,
		hub: H,
	}
	go s.readCircle()

	return s
}