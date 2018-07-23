package main

func main() {

}

// 基本的なRead
// buffer := make([]byte, 1024) // 1024バイトのメモリ領域確保
// size, err := r.Read(buffer)



// 補助関数

// 終端記号にあたるまですべてを読み込む
// buffer, err := ioutil.ReadAll(reader)

// 決まったバイト数だけ読み込む
// buffer := make([]byte, 4)
// size, err := ioutil.ReadFull(reader, buffer)

// すべてコピー
// writeSize, err := io.Copy(writer, reader)

// 指定したサイズだけコピー
// writeSize, err := io.CopyN(writer, reader, size)

// 予めコピーする量が決まっている、または何度もコピーするのでバッファを使いまわしたい
// buffer := make([]byte, 8 * 1024)
// io.CopyBuffer(writer, reader, buffer)
