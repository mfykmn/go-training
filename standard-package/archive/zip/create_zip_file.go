package main

import (
	"os"
	"archive/zip"
	"strings"
	"io"
)

// Goならわかるシステムプログラミング Q3.3の回答
func main() {
	file, err := os.Create("file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// ファイルの数だけ書き込み
  a, err := zipWriter.Create("a.txt")
  if err != nil {
  	panic(err)
  }
  io.Copy(a, strings.NewReader("1つめのファイルのテキストです"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("2つめのファイルのテキストです"))
}
