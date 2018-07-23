package network

type ConnMgr struct {
	conns map[int32]*Connector
}