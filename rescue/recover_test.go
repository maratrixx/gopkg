package rescue

import "testing"

func TestRunWithRecover(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fail()
		}
	}()

	RunWithRecover(func() {
		aa()
	})
}

func aa() {
	bb()
}

func bb() {
	panic("bb panic")
}
