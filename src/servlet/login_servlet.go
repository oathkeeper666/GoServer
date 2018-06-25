package servlet

import (
	// "logger"
)

type LoginSevlet struct {
	cmd uint8
	cmd2 uint8
	successor *Servlet
}

func (this *LoginSevlet) HandleMsg(cmd uint8) {
	
}

func (this *LoginSevlet) SetSuccessor(successor *Servlet) {
	this.successor = successor
}

func NewLoginServlet() *LoginSevlet {
	return &LoginSevlet {
		cmd: 0,
		cmd2: 100,
		successor: nil,
	}
}