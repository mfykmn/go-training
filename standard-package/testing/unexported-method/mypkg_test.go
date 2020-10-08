package unexported_test // テスト対象とは別のパッケージ

import (
	"log"
	"testing"

	"github.com/mfykmn/go-training/standard-package/testing/unexported-method"
)

func TestCounter(t *testing.T) {
	var c unexported.Counter
	c.Count()
	c.Count()
	c.Count()

	unexported.ExportCounterReset(&c)

	if 0 != c.Get() {
		log.Fatal("failed")
	} else {
		log.Println("success")
	}
}
