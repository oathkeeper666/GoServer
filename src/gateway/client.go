package network

import (
	"sync"
)

/*
	保存已经完成三次握手的session
*/
var sessionMap map[int64]*Session = make(map[int64]*Session)
var s_mutex *sync.Mutex = new(sync.Mutex)


/*
	已经完成登录操作的玩家
*/
var loginMap map[int64] *Session = make(map[int64]*Session)
var l_mutex *sync.Mutex = new(sync.Mutex)


func AddSession(s *Session) {
	s_mutex.Lock()
	defer s_mutex.Unlock()
	sessionMap[s.sid] = s
}

func RemoveSession(s *Session) {
	s_mutex.Lock()
	delete(sessionMap, s.sid)
	s_mutex.Unlock()

	// 如果客户端完成登录，移除登录的session
	RemoveLoginSession(s.pid)
}

/*
	通过sid获取session
*/
func GetSessionBySid(sid int64) *Session {
	s_mutex.Lock()
	delete(sessionMap, sid)
	return sessionMap[sid]
}

/*
	保存登录session
*/
func AddLoginSession(s *Session) {
	l_mutex.Lock()
	defer l_mutex.Unlock()
	loginMap[s.pid] = s
}

/*
	移除登录session
*/
func RemoveLoginSession(pid int64) {
	l_mutex.Lock()
	defer l_mutex.Unlock()
	delete(loginMap, pid)
}

/*
	通过玩家pid获取session
*/
func GetSessionByPid(pid int64) *Session {
	l_mutex.Lock()
	defer l_mutex.Unlock()
	return loginMap[pid]
}