package main

import (
	"fmt"
	"github.com/mfykmn/go-training/third-party-package/protoc-gen-go/addressbook"
)

func main() {
	p := tutorial.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "555-4321", Type: tutorial.Person_HOME},
		},
	}
	fmt.Printf("%#v", p)
}
