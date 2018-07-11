package main

import (
	"fmt"
	"log"

	"github.com/mafuyuk/go-training/third-party-package/gorilla/http"
)

// curl http://localhost:8080/v1/users
func main() {
	// Run App server
	services := &http.Handlers{}
	server := http.New(":8080", services)
	log.Println("Start server")
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed ListenAndServe. err: %v", err))
	}
}
