package syncx

import (
	"sync"
	"time"

	"github.com/imttx/golib/rescue"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) WaitTimeout(timeout time.Duration) bool {
	ch := make(chan struct{})
	go rescue.RunWithRecover(func() {
		defer close(ch)
		w.Wait()
	})

	select {
	case <-ch:
		return false
	case <-time.After(timeout):
		return true
	}
}
