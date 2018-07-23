package main

import (
	"log"
	"encoding/binary"
	"bytes"
	"fmt"
)

type Row struct {
	Key uint16 // 2バイト
	Val uint16
}

func (r Row) String() string {
	return fmt.Sprintf("(%s: %v)", string(r.Key), r.Val)
}

func main() {
	buf := bytes.NewBuffer([]byte{0x31, 0x97, 0x03, 0x1f})
	row := Row{}
	binary.Read(buf, binary.BigEndian, &row.Key)
	binary.Read(buf, binary.BigEndian, &row.Val)

	log.Println(row)
}
