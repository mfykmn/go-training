package main

import (
	"bufio"
	"io"

	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"strings"
)

func main() {
	var output io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		// 標準出力が端末なら出力時にバッファしない
		output = os.Stdout
	} else {
		output = bufio.NewWriter(os.Stdout)
	}

	for i := 0; i < 100; i++ {
		fmt.Fprintln(output, strings.Repeat("x", 100))
	}
	if _o, ok := output.(*bufio.Writer); ok {
		// bufio.Writerは最後にFlushを行う必要がある
		_o.Flush()
	}

}
