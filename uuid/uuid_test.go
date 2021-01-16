package uuid

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(UUID())
	}
}
func BenchmarkGenerateUUID(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			UUID()
		}
	})
}
