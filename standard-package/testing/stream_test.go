package testing

import (
	"testing"
	"os"
)

func streamBench(b *testing.B, file *os.File, f func(*os.File)) {
	b.ReportAllocs()
	for i:=0; i< b.N; i++ {
		f(file)
	}
}

//                                   ループが実行された回数 １ループごとの所要時間 １ループごとのアロケーションされたバイト数 １ループごとのアロケーション回数
// BenchmarkStream/MultiWriter-4             200000              5125 ns/op           38600 B/op         10 allocs/op
// BenchmarkStream/Seek-4                        50          28012993 ns/op         6317885 B/op       3951 allocs/op
func BenchmarkStream(b *testing.B) {
	benchCase := []struct {
		name string
		f func(*os.File)
	}{
		{"MultiWriter", multiWriter},
		{"Seek", seek},
	}

	for _, c := range benchCase {
		file, _ := os.Open("./sample.gif")
		b.Run(c.name, func(b *testing.B) { streamBench(b, file, c.f) })
		file.Close()
	}
}
