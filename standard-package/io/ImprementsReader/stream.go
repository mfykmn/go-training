package main

import (
	"strings"
	"io"
	"os"
)

var (
	computer = strings.NewReader("COMPUTER")
	system = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

// Goならわかるシステムプログラミングの回答 Q3.7
func main() {
	var stream io.Reader

	a := io.NewSectionReader(programming, 5, 1)
	c := io.LimitReader(computer, 1)
	s := io.LimitReader(system, 1)
	i := io.NewSectionReader(programming, 8, 1)
	pr, pw := io.Pipe()
	writer := io.MultiWriter(pw, pw)
	go io.CopyN(writer, i, 1)
	defer	pw.Close()

	stream = io.MultiReader(a, s, c, io.LimitReader(pr, 2))
	io.Copy(os.Stdout, stream)
}
