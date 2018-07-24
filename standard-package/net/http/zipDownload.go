package main

import (
	"net/http"
	"archive/zip"
	"io"
	"strings"
)

func main() {
	http.HandleFunc("/", zipHandler)
	http.ListenAndServe(":8080", nil)
}

func zipHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/zip")
	rw.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(rw)
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
