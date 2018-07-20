package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte("bytes.Buffer example\n"))
	buf.WriteString("bytes.Buffer example\n") //stringをそのまま使えるが、io.Writerのメソッドではないので他の構造体で使えない
	io.WriteString(&buf, "bytes.Buffer example\n")
	fmt.Println(buf.String())
}
