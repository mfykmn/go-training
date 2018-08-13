package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `1111,"ooooo","aaaaaa"
2222,"sssss","bbbbbb"
3,"ddddddd","cccccc"
`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}
