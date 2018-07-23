# 基本的なRead
```go
buffer := make([]byte, 1024) // 1024バイトのメモリ領域確保
size, err := r.Read(buffer)
```


# 補助関数
* 終端記号にあたるまですべてを読み込む
```go
buffer, err := ioutil.ReadAll(reader)
```

* 決まったバイト数だけ読み込む
```go
buffer := make([]byte, 4)
size, err := ioutil.ReadFull(reader, buffer)
```

* すべてコピー
```go
writeSize, err := io.Copy(writer, reader)
```

* 指定したサイズだけコピー
```go
writeSize, err := io.CopyN(writer, reader, size)
```

* 予めコピーする量が決まっている、または何度もコピーするのでバッファを使いまわしたい
```go
buffer := make([]byte, 8 * 1024)
io.CopyBuffer(writer, reader, buffer)
```

# インターフェース
* Reader を readCloserにできる　テストなどで使う
```go
var reader io.Reader = strings.NewReader("テストデータ")

var readCloser io.ReadCloser = ioutil.NopCloser(reader)
readCloser.Close()
```

* 個別のreaderとwriterをつなげてreadWriterを作れる
```go
var readWriter io.ReadWriter = bufio.NewReadWriter(reader, writer)
```

# 必要な部分だけ切り出す
* 先頭の16バイトしか読み込めないようにする
```go
r := io.LimitReader(reader, 16)
```
