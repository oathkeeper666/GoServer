package network

import (
	"net"
	"consts"
	"time"
)

type Connector struct {
	id int32
	protocol string
	address string
	conn *net.TCPConn
	isConnected bool
}

var i int32 = 0

func NewConnector(protocol string, address string) *Connector {
	i++
	return &Connector {
		id: i,
		protocol: protocol,
		address: address,
		conn: nil,
		isConnected: false,
	}
}

func (this *Connector) Connect() bool {
	conn, err := net.Dial(this.protocol, this.address)
	if err == nil {
		this.conn = conn.(*net.TCPConn)
		this.conn.SetReadBuffer(consts.CONN_WRITER_BUFFF_SIZE)
		this.conn.SetWriteBuffer(consts.CONN_READ_BUFF_SIZE)
		this.isConnected = true
	}

	return this.isConnected
}

func (this *Connector) checkConnect() {
	if !this.isConnected {
		for !this.Connect() {
			time.Sleep(time.Second)
		}
	}
}