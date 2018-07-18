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


	// パースデモ
	const maxBandWidth = "10MB"
	conf := &Conf{}
	conf.Parse(maxBandWidth)
	fmt.Printf("Pased %s to %d", "10MB", conf.maxBandWidth)

}

type Conf struct {
	maxBandWidth uint64
}

func (c *Conf)Parse(maxBandWidth string) {
	if bw, err := humanize.ParseBytes(maxBandWidth); err != nil {
		fmt.Println("Can not parse -max-bandwidth", err)
		os.Exit(1)
	} else {
		c.maxBandWidth = bw
	}
}
