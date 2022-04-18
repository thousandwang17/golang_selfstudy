package test

import (
	"testing"
)

type Item2 struct {
	id  int
	val [4096]byte
}

func generateItems(n int) []*Item2 {
	items := make([]*Item2, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item2{id: i})
	}
	return items
}

func BenchmarkForPointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangePointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}

/*
go test -bench=Pointer$ .
goos: darwin
goarch: arm64
pkg: forRange/test
BenchmarkForPointer-8            1539914               707.3 ns/op
BenchmarkRangePointer-8          1671140               698.4 ns/op
*/
