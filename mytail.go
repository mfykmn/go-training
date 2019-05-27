package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

var reg = regexp.MustCompile("\r\n|\n\r|\n|\r")

func main() {
	var n *int = flag.Int("n", 10, "num")
	flag.Parse()

	file, err := os.Open(os.Args[3])
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed open file.\n")
		os.Exit(1)
		return
	}
	defer file.Close()


	buf := make([]byte, 1024)
	if _, err := file.Read(buf); err != nil {
		fmt.Fprint(os.Stderr, "Failed read file.\n")
		os.Exit(1)
		return
	}

	spliterd := reg.Split(string(buf), -1)
	if (len(spliterd) - *n) < 0 {
		for _, v := range spliterd {
			fmt.Println(v)
		}
	} else {
		for _, v := range spliterd[len(spliterd) - *n:] {
			fmt.Println(v)
		}
	}

}
