package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	abspath, _ := filepath.Abs("testfile")
	fmt.Println(abspath)
}
