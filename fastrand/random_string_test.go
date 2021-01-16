package fastrand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	rs := RandString(30)
	t.Logf("%+v\n", rs)
	assert.Equal(t, 30, len(rs))
}

func BenchmarkRandString(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			RandString(30)
		}
	})
}
