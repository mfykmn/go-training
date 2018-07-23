package main

import (
	"io"
	"strings"
	"io/ioutil"
)

func main() {
	var reader io.Reader = strings.NewReader("テストデータ")
	var readCloser io.ReadCloser = ioutil.NopCloser(reader)
	readCloser.Close()
}
