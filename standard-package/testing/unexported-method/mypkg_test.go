package unexported_test // テスト対象とは別のパッケージ

import (
	"testing"
	"log"

	"github.com/mafuyuk/go-training/standard-package/testing/unexported-method"
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