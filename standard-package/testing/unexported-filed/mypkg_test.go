package unexported_test // テスト対象とは別のパッケージ

import (
	"testing"

	"github.com/mafuyuk/go-training/standard-package/testing/unexported-filed"
)

func TestCounter(t *testing.T) {
	var c unexported.Counter
	c.ExportSetN(1)
}