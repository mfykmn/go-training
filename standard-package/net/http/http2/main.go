package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

// ブラウザから https://localhost:8080/ にアクセス
func main() {
	certFile, _ := filepath.Abs("server.crt")
	keyFile, _ := filepath.Abs("server.key")
	fmt.Println(certFile)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	})

	err := http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
