package main

import (
	"os"
	"fmt"

	"github.com/dustin/go-humanize"
)

func main() {
	name := os.Args[1]
	s, _ := os.Stat(name)
	fmt.Printf(
		"%s: %s\n",
		name,
		humanize.Bytes(uint64(s.Size())),
	)
}
