package io

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"strings"
	"io"
)

var in1 = strings.NewReader(`{"name":"mafuyuk","country":"japan"}`)
var in2 = strings.NewReader(`{"name":"mafuyuk","country":"japan"}`)

type Foo struct {
	Name string `json:"name"`
	Country string `json:"country"`
}
var u1 Foo
var u2 Foo

func BenchmarkJSON_ReadAll(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := ioutil.ReadAll(in1)
		if err != nil {
			b.Fatal(err)
		}
		err = json.Unmarshal(data, &u1)
		if err != nil {
			b.Fatal(err)
		}
		in1.Seek(0, io.SeekStart)
	}
}

func BenchmarkJSON_Decoder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.NewDecoder(in2).Decode(&u2)
		if err != nil {
			b.Fatal(err)
		}
		in2.Seek(0, io.SeekStart)
	}
}

//$ cd io
//$ go test -bench . -benchmem
//goos: darwin
//goarch: amd64
//pkg: github.com/mafuyuk/go-training/io
//BenchmarkJSON_ReadAll-4          1000000              1771 ns/op            2472 B/op          7 allocs/op
//BenchmarkJSON_Decoder-4          1000000              1345 ns/op            1024 B/op          6 allocs/op
//PASS
//ok      github.com/mafuyuk/go-training/io       4.098s

// テスト名 実行した回数 １回あたりの実行に掛かった時間(ns/op)  １回あたりのアロケーションで確保した容量(B/op)  1回あたりのアロケーション回数(allocs/op)

