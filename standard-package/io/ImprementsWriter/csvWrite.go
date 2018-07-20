package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"a", "b", "c"})
	writer.Flush()
	writer.Write([]string{"d", "e", "f"})
	writer.Flush()
	writer.Write([]string{"g", "h", "i"})
	writer.Flush()
}
