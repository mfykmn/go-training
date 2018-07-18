package main

import (
	"os/exec"
	"fmt"
)


func main() {
	if out, err := exec.Command("uname", "-s").Output(); err != nil {
		switch e := err.(type) {
		case *exec.Error:
			// コマンドは実行したものの異常終了した
			fmt.Println("exec.Error", err, e)
		case *exec.ExitError:
			// コマンドは実行したが異常終了
			fmt.Println("exec.ExitError", err, e)
		}
	} else {
		fmt.Println(string(out))
	}
}
