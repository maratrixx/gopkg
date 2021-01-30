package syncx

import (
	"sync"
	"time"

	"github.com/imttx/gopkg/rescue"
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
	default:
		timer := time.NewTimer(timeout)
		select {
		case <-ch:
			timer.Stop()
			return false
		case <-timer.C:
			return true
		}
	}
}
