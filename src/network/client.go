package network

import (
	"net"
	"bytes"
	"fmt"
)

type TransHeader struct {
	cmd uint8
	other string
}

var clientMap map[string]*Client

type Client struct {
	conn net.Conn
	remoteIp string
	readBuf *bytes.Buffer
	writeBuf *bytes.Buffer
}

func NewClient(conn net.Conn) (*Client) {
	if clientMap == nil {
		clientMap = make(map[string]*Client)
	}
	ip := conn.RemoteAddr().String()
	clientMap[ip] = &Client {
		conn: conn,
		remoteIp: conn.RemoteAddr().String(),
		readBuf: bytes.NewBuffer(make([]byte, 256)),
		writeBuf: bytes.NewBuffer(make([]byte, 256)),
	}

	return clientMap[ip]
}

func readCircle(c *Client) {
	for {
		b := make([]byte, 256)
		_, err := c.conn.Read(b)
		if err == nil {
			c.readBuf.Write(b)	
		} else {
			fmt.Printf("read from %s error, error is %v\n", c.remoteIp, err)
		}
	}
}

func writeCircle(c *Client) {
	for {
		if c.writeBuf.Len() != 0 {
			b := make([]byte, 256)
			c.writeBuf.Read(b)
			_, err := c.conn.Write(b)
			if err != nil {
				c.writeBuf.Write(b)
				fmt.Printf("write to %s error, err is %v\n", c.remoteIp, err)
			}
		}
	}
}

func HandleConnection(c *Client) {
	go readCircle(c)
	go writeCircle(c)
}