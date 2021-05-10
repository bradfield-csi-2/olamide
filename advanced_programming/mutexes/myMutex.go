package myMutex

import (
	"sync/atomic"
)

type MyMutex struct {
	counter uint32
}

func (m *MyMutex) lock() {

	for {
		if atomic.LoadUint32(&m.counter) == 0 {
			atomic.AddUint32(&m.counter, 1)
			break
		}
	}

}

func (m *MyMutex) unlock() {
	if atomic.LoadUint32(&m.counter) == 1 {
		atomic.StoreUint32(&m.counter, 0)
	} else {
		panic("Error, unlocking multiple times")
	}
}
