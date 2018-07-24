package util

import (
	"time"
	"sync"
	"logger"
)

var (
	operataAllocPool = sync.Pool {
		New: func () interface{} {
			return &Operation {}
		},
	}
)

type Operation struct {
	name string
	startTime time.Time
}

func StartOperation(name string) *Operation {
	op := operataAllocPool.Get().(*Operation)
	op.name = name
	op.startTime = time.Now()
	return op
}

func (op *Operation) Finish(warnThreshold time.Duration) {
	takeTime := time.Now().Sub(op.startTime)
	if takeTime >= warnThreshold {
		logger.WRITE_WARNING("operation %s takes %s", op.name, takeTime)
	}
	operataAllocPool.Put(op)
}
