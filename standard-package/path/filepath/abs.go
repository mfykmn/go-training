package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	abspath, _ := filepath.Abs("testfile")
	fmt.Println(abspath)
}
