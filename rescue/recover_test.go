package rescue

import "testing"

func TestRunWithRecover(t *testing.T) {
	RunWithRecover(func() {
		aa()
	})
}

func aa() {
	bb()
}

func bb() {
	cc()
}

func cc() {
	panic("cc panic")
}
