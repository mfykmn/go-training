package main

import (
	"fmt"

	"github.com/mfykmn/go-training/third-party-package/xid/xid"
)

func main() {
	idCreator := xid.New("user_")

	fmt.Println(idCreator.Create())
	fmt.Println(idCreator.Create())
}
