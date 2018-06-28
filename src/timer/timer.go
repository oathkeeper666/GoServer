package timer

import (
	"time"
)



// 回调函数
type CallBackFunc struct {
	id int8
	f func(interface {})
	arg interface  {}
}

var func_chan chan *CallBackFunc = make 
var g_time time.Time;	// 服务器时间
var g_func_id int8

// 注册一个定时处理的函数
func RegisterFunc(f func(interface {}), secs uint32, ) int8 {
	
}
