package main

import (
	"os/exec"
	"fmt"
)


func main() {
	out, _ := exec.Command("uname", "-s").Output()
	fmt.Println(string(out))
}
