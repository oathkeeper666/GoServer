package network

import (
	"sync"
)

const (
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