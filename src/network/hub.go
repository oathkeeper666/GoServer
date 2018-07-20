package network

import (
	"logger"
)

type Hub struct {
	// registered session
	sess map[*Session]bool
	// broadcast message
	broadcast chan []byte
	// register
	register chan *Session
	// unregister
	unregister chan *Session
}

var H = NewHub()

func NewHub() *Hub {
	return &Hub {
		sess: make(map[*Session]bool),
		broadcast: make(chan []byte, 10),
		register: make(chan *Session, 10),
		unregister: make(chan *Session, 10),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			h.sess[s] = true
		case s := <-h.unregister:
			s.Close()
			delete(h.sess, s)
		case message := <-h.broadcast:
			logger.WRITE_DEBUG(string(message))
		}
	}
}
