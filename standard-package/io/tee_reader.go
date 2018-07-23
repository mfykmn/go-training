package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)

	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)

	// けどバッファに残っている
	fmt.Println(buffer.String())
}
