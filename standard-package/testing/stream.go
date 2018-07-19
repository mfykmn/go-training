package testing

import (
	"os"
	"image"
	"image/gif"
	"io"
	"bytes"
)

func multiWriter(in *os.File) {
	buf := new(bytes.Buffer)
	gifDecodeTarget := new(bytes.Buffer)

	w := io.MultiWriter(buf, gifDecodeTarget)
	io.Copy(w, in)

	image.Decode(buf)
	gif.DecodeAll(gifDecodeTarget)
}

func seek(in *os.File) {
	image.Decode(in)

	in.Seek(0, 0) // インデックスを先頭に戻す
	gif.DecodeAll(in)
}
