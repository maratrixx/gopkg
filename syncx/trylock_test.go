package syncx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMutex_TryLock(t *testing.T) {
	m := NewMutex()

	ok := m.TryLock(100 * time.Millisecond)
	assert.Equal(t, true, ok, "it should lock succ, but fail")

	ok = m.TryLock(100 * time.Millisecond)
	assert.Equal(t, false, ok, "it should lock fail, but succ")

	m.Unlock()
	ok = m.TryLock(100 * time.Millisecond)
	assert.Equal(t, true, ok, "it should lock succ, but fail")
}

func BenchmarkTrylock(b *testing.B) {
	var number int
	lock := NewMutex()
	for i := 0; i < b.N; i++ {
		go func() {
			defer lock.Unlock()
			lock.TryLock(time.Microsecond)
			number++
		}()
	}
}
