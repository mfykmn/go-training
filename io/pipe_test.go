package io

import (
	"io"
	"bytes"
	"fmt"
)

func ExamplePipeNG() {
	print := func(r io.Reader) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		fmt.Print(buf.String())
	}

	b := new(bytes.Buffer) //不要なアロケート
	fmt.Fprint(b, "some text to be read\n")
	print(b)

	// Output:
	// some text to be read
}
