package hack

import (
	"reflect"
	"testing"
)

func TestToSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantB []byte
	}{
		{args: args{s: "hello"}, wantB: []byte("hello")},
		{args: args{s: ""}, wantB: []byte(nil)},
		{args: args{s: ""}, wantB: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := ToSlice(tt.args.s); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ToSlice() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{args: args{b: []byte("hello")}, wantS: "hello"},
		{args: args{b: []byte("")}, wantS: ""},
		{args: args{b: []byte(nil)}, wantS: ""},
		{args: args{b: nil}, wantS: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := ToString(tt.args.b); gotS != tt.wantS {
				t.Errorf("ToString() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

var sdata = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var bdata = []byte(sdata)

func BenchmarkToStringStd(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = string(bdata)
		}
	})
}
func BenchmarkToString(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = ToString(bdata)
		}
	})
}

func BenchmarkToSliceStd(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = []byte(sdata)
		}
	})
}
func BenchmarkToSlice(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = ToSlice(sdata)
		}
	})
}

// benchmark
// go test -benchmem -run=^$ -bench ^(BenchmarkToStringStd|BenchmarkToString|BenchmarkToSliceStd|BenchmarkToSlice)$ github.com/imttx/gopkg/hack -v -count=1

// goos: darwin
// goarch: amd64
// pkg: github.com/imttx/gopkg/hack
// BenchmarkToStringStd
// BenchmarkToStringStd-8   	66936225	        18.0 ns/op	      64 B/op	       1 allocs/op
// BenchmarkToString
// BenchmarkToString-8      	1000000000	         0.341 ns/op	       0 B/op	       0 allocs/op
// BenchmarkToSliceStd
// BenchmarkToSliceStd-8    	58576948	        21.4 ns/op	      64 B/op	       1 allocs/op
// BenchmarkToSlice
// BenchmarkToSlice-8       	1000000000	         0.390 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/imttx/gopkg/hack	3.320s
