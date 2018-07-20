package network

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
		sess: make(map[*Session]bool)
		broadcast: make(chan []byte),
		register: make(chan *Session),
		unregister: make(chan *Session),
	}
}

func (h *Hub) run() {
	for {
		select {
		case s := <-h.register:
			h.sess[s] = true
		case s := <-h.unregister:
			s.Close()
			delete(h.sess, s)
		case message := <-h.broadcast:

		}
	}
}