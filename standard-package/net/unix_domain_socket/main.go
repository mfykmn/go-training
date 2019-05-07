package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, _ := net.Listen("unix", "socket")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go func() {
			fmt.Printf("Accept:%v\n", conn.RemoteAddr())
			request, _ := http.ReadRequest(bufio.NewReader(conn))
			dump, _ := httputil.DumpRequest(request, true)
			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello I am Response.")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
