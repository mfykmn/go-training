package unexported_test // テスト対象とは別のパッケージ

import (
	"testing"

	"github.com/mfykmn/go-training/standard-package/testing/unexported"
)

func TestMypkg(t *testing.T) {
	// maxValueの代わりにExportMaxValueを参照する
	if doSomething() > unexported.ExportMaxValue {
		t.Error("Error")
	}
}

func doSomething() int {
	return 99
}
