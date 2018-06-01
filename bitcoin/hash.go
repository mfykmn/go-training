package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	fmt.Print("aaaa")
	str := "go love"
	fmt.Println(sha256.Sum256([]byte(str)))
	fmt.Println(sha512.Sum512([]byte(str)))

	salt := []byte("this is salt")
	hash := pbkdf2.Key([]byte(str), salt,4096,32, sha256.New)
	fmt.Println(string(hash))
}
