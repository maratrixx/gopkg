package syncx

import (
	"testing"
	"time"

	"github.com/imttx/golib/rescue"
)

func TestWaitGroupWrapper_Wait(t *testing.T) {
	var wg WaitGroupWrapper
	wg.Add(1)

	go rescue.RunWithRecover(func() {
		time.Sleep(1 * time.Second)
		wg.Done()
	})

	t0 := time.Now()
	wg.Wait()
	t.Logf("wait for %s", time.Since(t0))
}

func TestWaitGroupWrapper_Waitout(t *testing.T) {
	var wg WaitGroupWrapper
	wg.Add(1)

	go rescue.RunWithRecover(func() {
		time.Sleep(3 * time.Second)
		wg.Done()
	})

	t0 := time.Now()
	wg.WaitTimeout(300 * time.Millisecond)
	t.Logf("wait for %s", time.Since(t0))
}
