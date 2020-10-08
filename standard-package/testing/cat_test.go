package main

import (
	"fmt"
	"testing"
)

func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")
	}
	return s
}

func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}

func BenchmarkCat3(b *testing.B)     { bench(b, 3, cat) }
func BenchmarkBuf3(b *testing.B)     { bench(b, 3, buf) }
func BenchmarkCat100(b *testing.B)   { bench(b, 100, cat) }
func BenchmarkBuf100(b *testing.B)   { bench(b, 100, buf) }
func BenchmarkCat10000(b *testing.B) { bench(b, 10000, cat) }
func BenchmarkBuf10000(b *testing.B) { bench(b, 10000, buf) }

//$ go test -bench .
//goos: darwin
//goarch: amd64
//pkg: github.com/mfykmn/go-training/standard-package/testing

//                ループが実行された回数 １ループごとの所要時間 １ループごとのアロケーションされたバイト数 １ループごとのアロケーション回数
//BenchmarkCat3-4           10000000           127 ns/op              54 B/op                             3 allocs/op
//BenchmarkBuf3-4           10000000           137 ns/op             163 B/op                             3 allocs/op
//BenchmarkCat100-4           300000          4284 ns/op            7520 B/op                           100 allocs/op
//BenchmarkBuf100-4          1000000          1198 ns/op            2160 B/op                             4 allocs/op
//BenchmarkCat10000-4            200       6383734 ns/op        53327792 B/op                         10000 allocs/op
//BenchmarkBuf10000-4          10000        102985 ns/op          211424 B/op                            11 allocs/op

// サブベンチマーク
// Table Drivenなベンチマークを記述できる
func BenchmarkConcatenate(b *testing.B) {
	benchCase := []struct {
		name string
		n    int
		f    func(...string) string
	}{
		{"Cat", 3, cat},
		{"Buf", 3, buf},
		{"Cat", 100, cat},
		{"Buf", 100, buf},
		{"Cat", 10000, cat},
		{"Buf", 10000, buf},
	}
	for _, c := range benchCase {
		b.Run(fmt.Sprintf("%s%d", c.name, c.n),
			func(b *testing.B) { bench(b, c.n, c.f) })
	}
}
