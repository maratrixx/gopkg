package syncx

import (
	"testing"
	"time"

	"github.com/imttx/gopkg/rescue"
	"github.com/stretchr/testify/assert"
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
		time.Sleep(time.Second)
		wg.Done()
	})

	t0 := time.Now()
	res := wg.WaitTimeout(100 * time.Millisecond)
	assert.Equal(t, true, res)
	t.Logf("wait_timeout for %s", time.Since(t0))
}
