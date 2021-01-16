package fastrand

import (
	"math/rand"
	"testing"
	_ "unsafe"

	"github.com/stretchr/testify/assert"
)

func TestUint32(t *testing.T) {
	v := Uint32()
	t.Logf("%+v\n", v)
}

func TestUint64(t *testing.T) {
	v := Uint64()
	t.Logf("%+v\n", v)
}

func TestUint32N(t *testing.T) {
	v := Uint32N(10)
	t.Logf("%+v\n", v)
	assert.Equal(t, true, v < 10)
}

func TestUint64N(t *testing.T) {
	v := Uint64N(10)
	t.Logf("%+v\n", v)
	assert.Equal(t, true, v < 10)
}

func BenchmarkFastranUint32(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Uint32()
		}
	})
}

func BenchmarkStdUint32(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			rand.Int31()
		}
	})
}

func BenchmarkFastranUint64(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Uint64()
		}
	})
}

func BenchmarkStdUint64(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			rand.Int63()
		}
	})
}

func BenchmarkFastrandUint32N(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Uint32N(100)
		}
	})
}

func BenchmarkStdUint32N(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			rand.Int31n(100)
		}
	})
}

// go test -benchmem -run=^$ -bench ^(BenchmarkFastranUint32|BenchmarkStdUint32|BenchmarkFastranUint64|BenchmarkStdUint64|BenchmarkFastrandUint32N|BenchmarkStdUint32N)$ github.com/imttx/golib/fastrand -v -count=1
//
// goos: darwin
// goarch: amd64
// pkg: github.com/imttx/golib/fastrand
// BenchmarkFastranUint32
// BenchmarkFastranUint32-8     	1000000000	         0.891 ns/op	       0 B/op	       0 allocs/op
// BenchmarkStdUint32
// BenchmarkStdUint32-8         	14682615	        81.1 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFastranUint64
// BenchmarkFastranUint64-8     	510861272	         2.14 ns/op	       0 B/op	       0 allocs/op
// BenchmarkStdUint64
// BenchmarkStdUint64-8         	14689384	        81.5 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFastrandUint32N
// BenchmarkFastrandUint32N-8   	693901076	         1.77 ns/op	       0 B/op	       0 allocs/op
// BenchmarkStdUint32N
// BenchmarkStdUint32N-8        	12233980	        96.4 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/imttx/golib/fastrand	7.567s
