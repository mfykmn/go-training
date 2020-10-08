package main

import (
	"io"
)

type MyReader struct {
	i int
}

func (mr MyReader) Read(b []byte) (int, error) {
	if len(b) < mr.i {
		return 0, io.EOF
	}

	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}

	mr.i += len(b)
	return len(b), nil
}
