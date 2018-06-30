package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	str := "go love"

	fmt.Printf("SHA-256 : %x\n", sha256.Sum256([]byte(str)))
	fmt.Printf("SHA-512 : %x\n", sha512.Sum512([]byte(str)))

	salt := []byte("this is salt")
	hash := pbkdf2.Key([]byte(str), salt, 4096, 32, sha256.New)
	fmt.Println(string(hash))
}
