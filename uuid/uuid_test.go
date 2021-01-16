package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	uuid := UUID()
	t.Logf("%v\n", uuid)

	assert.Equal(t, 36, len(uuid))
}

func BenchmarkGenerateUUID(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			UUID()
		}
	})
}
