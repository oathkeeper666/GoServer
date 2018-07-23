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

func (this *LoginSevlet) HandleMsg(cmd int32, buffer []byte) ([]byte) {
	if this.cmd >= cmd && this.cmd2 < cmd {
		return nil
	}
	if this.successor != nil {
		return this.successor.HandleMsg(cmd, buffer)
	}

	return nil
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