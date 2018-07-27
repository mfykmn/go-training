package main

import (
	"net"
	"net/http"
	"bufio"
	"net/http/httputil"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
