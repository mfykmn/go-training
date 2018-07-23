package main

import (
	"strings"
	"io"
	"os"
)

func main() {
	reader := strings.NewReader("Example of io.SectionReader\n") // ReaderAtを満たしている
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
