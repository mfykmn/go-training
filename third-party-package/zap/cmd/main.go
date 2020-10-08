package main

import (
	"log"

	"github.com/mfykmn/go-training/third-party-package/zap"
)

//  go run cmd/main.go 1> stdout.log 2> stderr.log
func main() {
	zap.New(zap.Debug)
	zap.Infof("dummy", "標準出力")
	zap.Errorf("dummy", "標準エラー出力")
	zap.Fatalf("dummy", "標準エラー出力")

	log.Println("これは標準エラー出力にでる")
}
