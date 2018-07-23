package network

import (
	"logger"
	"consts"
	"time"
)

type Hub struct {
	// registered session
	sess map[int64]*Session
	// broadcast message
	broadcast chan []byte
	// register
	register chan *Session
	// unregister
	unregister chan *Session
	// receive peer package queue
	receiveMsgQueue chan *Package
	// for receiveMsgQueue
	ticker <-chan time.Time
}

var H = NewHub()

func NewHub() *Hub {
	return &Hub {
		sess: make(map[int64]*Session),
		broadcast: make(chan []byte, consts.HUB_BROADCAST_QUEUE_SIZE),
		register: make(chan *Session, consts.HUB_ACCEPT_SIZE),
		unregister: make(chan *Session, consts.HUB_UNACCEPT_SIZE),
		receiveMsgQueue: make(chan *Package, consts.HUB_SEND_QUEUE_SIZE),
		ticker: time.Tick(consts.HUB_SERVICE_TIME),
	}
}

/*
	处理保存下来的未处理的协议包
*/
func (h *Hub) flushPendingPackets() {
	for pack := range h.receiveMsgQueue {
		h.sess[pack.sid].ProcessData(pack)
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			h.sess[s.sid] = s
		case s := <-h.unregister:
			s.Close()
			delete(h.sess, s.sid)
		case message := <-h.broadcast:
			logger.WRITE_DEBUG(string(message))
		case <-h.ticker:
			h.flushPendingPackets()
		}
	}
}
