package servlet

type Servlet interface {
	HandleMsg(cmd uint8)
	SetSuccessor(successor *Servlet)
}