package servlet

type Servlet interface {
	HandleMsg(cmd int32, buffer []byte) ([]byte)
	SetSuccessor(successor Servlet)
}