package testing

import (
	"log"
	"os"
	"testing"
)

func streamBench(b *testing.B, f func(*os.File)) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		file, err := os.Open("./sample.gif")
		if err != nil {
			log.Println("Failed file open")
			return
		}

		f(file)
		file.Close()
	}
}

// $ go test -bench BenchmarkStream
//                                   ループが実行された回数 １ループごとの所要時間 １ループごとのアロケーションされたバイト数 １ループごとのアロケーション回数
// BenchmarkStream/MultiWriter-4                 20          54441881 ns/op        16391394 B/op       7903 allocs/op
// BenchmarkStream/Seek-4                        20          55851878 ns/op        12626858 B/op       7890 allocs/op
func BenchmarkStream(b *testing.B) {
	benchCase := []struct {
		name string
		f    func(*os.File)
	}{
		{"MultiWriter", multiWriter},
		{"Seek", seek},
	}

	for _, c := range benchCase {
		b.Run(c.name, func(b *testing.B) { streamBench(b, c.f) })
	}
}
