package syncx

import "time"

type Mutex struct {
	c chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{c: make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.c <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.c
}

func (m *Mutex) TryLock(timtout time.Duration) bool {
	select {
	case m.c <- struct{}{}:
		return true
	default:
		timer := time.NewTimer(timtout)
		select {
		case m.c <- struct{}{}:
			timer.Stop()
			return true
		case <-timer.C:
			return false
		}
	}
}
