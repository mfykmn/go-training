package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// "/"でパス文列を結合しない
	// windowsだとパスが\だから動きがおかしくなる
	dir := filepath.Join(u.HomeDir, ".config", "myapp")
	fmt.Println(dir) // $HOMW/.config/myapp

	// path/filepathパッケージは物理パスを操作するために使う
	// pathパッケージはhttpやftpなどの論理パスを操作するために使う
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
