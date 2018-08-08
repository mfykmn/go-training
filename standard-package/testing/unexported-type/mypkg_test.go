package unexported_test // テスト対象とは別のパッケージ

import (
	"testing"
	"encoding/json"
	"log"

	"github.com/mafuyuk/go-training/standard-package/testing/unexported-type"

)

func TestClient(t *testing.T) {
	response := &unexported.ExportResponse{
		Vaue: "testtest",
	}
	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}