package unexported_test // テスト対象とは別のパッケージ

import (
	"testing"

	"github.com/mfykmn/go-training/standard-package/testing/unexported-val"
)

func TestClient(t *testing.T) {
	// SetBaseURLで返ってきた関数をdeferで呼び出す
	defer unexported.SetBaseURL("http://localshot:8080")()

	// 以下にテストコード
}
