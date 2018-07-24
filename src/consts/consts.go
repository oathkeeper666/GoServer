package consts

import (
	"time"
)

const (
	SESSION_READ_MSG_SIZE = 1024
	SESSION_WRITE_MSG_SIZE = 1024
	SESSION_HEART_BEATE_INTERVAL = 60 * time.Second

	HUB_BROADCAST_QUEUE_SIZE = 1024 * 1024
	HUB_ACCEPT_SIZE = 10
	HUB_UNACCEPT_SIZE = 10
	HUB_SEND_QUEUE_SIZE = 1024 * 1024
	HUB_SERVICE_TIME = time.Millisecond * 5

	CONN_WRITER_BUFFF_SIZE = 1024 * 1024
	CONN_READ_BUFF_SIZE = 1024 * 1024

	WRITE_WAIT_TIME = 10 * time.Second
	PONG_WAIT = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	PIND_PERIOD = (PONG_WAIT * 9) / 10
	// Maximum message size allowed from peer.
	MAX_MESSAGE_SIZE = 512
)