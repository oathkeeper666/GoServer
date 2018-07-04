package servlet

type Servlet interface {
	HandleMsg(sid int64, cmd int32, buffer []byte)
	SetSuccessor(successor Servlet)
}