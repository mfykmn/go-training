package testing

import (
	"bytes"
	"image"
	"image/gif"
	"io"
	"os"
)

func multiWriter(in *os.File) {
	buf := new(bytes.Buffer)
	gifDecodeTarget := new(bytes.Buffer)

	w := io.MultiWriter(buf, gifDecodeTarget)
	_, err := io.Copy(w, in)
	if err != nil {
		panic(err)
	}

	_, _, err = image.Decode(buf)
	if err != nil {
		panic(err)
	}

	_, err = gif.DecodeAll(gifDecodeTarget)
	if err != nil {
		panic(err)
	}
}

func seek(in *os.File) {
	_, _, err := image.Decode(in)
	if err != nil {
		panic(err)
	}

	in.Seek(0, 0) // インデックスを先頭に戻す
	_, err = gif.DecodeAll(in)
	if err != nil {
		panic(err)
	}
}
