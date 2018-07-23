package main

import (
	"bufio"
	"strings"
	"fmt"
)

var source = `１行目
２行目 です
３行目`

func main () {
	useBufioReader()
	useBufioScanner()
}

func useBufioReader() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err != nil {
			break
		}
	}
}

// 分割文字が削除されることに注意
func useBufioScanner() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
