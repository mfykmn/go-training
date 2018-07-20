package main

import (
"net"
"io"
"os"
"net/http"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("GET", "http://ascii.jp", nil)
	req.Write(conn)

	io.Copy(os.Stdout, conn)

}
