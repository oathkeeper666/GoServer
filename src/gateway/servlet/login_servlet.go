package servlet

import (
	// "logger"
	"servlet"
)

type LoginSevlet struct {
	cmd int32
	cmd2 int32
	successor servlet.Servlet
}

func (this *LoginSevlet) HandleMsg(sid int64, cmd int32, buffer []byte) {
	if this.cmd >= cmd && this.cmd2 < cmd {

	} else {
		if this.successor != nil {
			this.successor.HandleMsg(sid, cmd, buffer)
		}
	}
}

func (this *LoginSevlet) SetSuccessor(successor servlet.Servlet) {
	this.successor = successor
}

func NewLoginServlet() *LoginSevlet {
	return &LoginSevlet {
		cmd: 0,
		cmd2: 100,
		successor: nil,
	}
}