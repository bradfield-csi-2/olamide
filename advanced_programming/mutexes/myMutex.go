package myMutex

import (
	"sync/atomic"
)

type MyMutex struct {
	counter uint32
}

func (m *MyMutex) lock() {

	for {
		if atomic.CompareAndSwapUint32(&m.counter, 0, 1) == true {
			return
		}
	}
}

func (m *MyMutex) unlock() {
	if atomic.CompareAndSwapUint32(&m.counter, 1, 0) == true {
		return
	}
}
