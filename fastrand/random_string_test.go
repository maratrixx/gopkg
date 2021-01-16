package fastrand

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandString(30))
	}
}

func BenchmarkRandString(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			RandString(30)
		}
	})
}
