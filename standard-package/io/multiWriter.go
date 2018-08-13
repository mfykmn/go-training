package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// $ echo -n foo | go run multiWriter.go
// Wrote 3 bytes to /var/folders/s4/29qs3_hd6dndvz58b7h2j12c0000gp/T/tmp650842643
// SHA256: 2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
func main() {
	// テンポラリファイルを開く
	tmp, err := ioutil.TempFile(os.TempDir(), "tmp")
	if err != nil {
		fmt.Println("error create temp file")
		return

	}
	defer tmp.Close()

	// SHA256計算用
	hash := sha256.New()

	// 両方に書き込むためのio.MultiWriter
	w := io.MultiWriter(tmp, hash)

	// io.Copyで標準入力からMultiWriterへコピー
	written, err := io.Copy(w, os.Stdin)
	if err != nil {
		fmt.Println("error io.Copy")
		return
	}

	fmt.Printf("Wrote %d bytes to %s\nSHA256: %x\n",
		written,       // 書き込まれたバイト数
		tmp.Name(),    // テンポラリファイル名
		hash.Sum(nil), // ハッシュ値
	)
}
