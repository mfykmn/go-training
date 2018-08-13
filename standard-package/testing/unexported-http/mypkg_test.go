package unexported_test // テスト対象とは別のパッケージ

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"encoding/json"
	"github.com/mafuyuk/go-training/standard-package/testing/unexported-http"
)

func TestGet(t *testing.T) {
	cases := map[string]struct {
		n        int
		hasError bool
	}{
		"100": {n: 100},
		"200": {n: 200},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			var requested bool
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requested = true
				if r.FormValue("n") != strconv.Itoa(tc.n) {
					t.Errorf("param n want %s got %d", r.FormValue("n"), tc.n)
				}
				resp := &unexported.ExportGetResponse{
					Value: "hoge",
				}
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					t.Fatal("unexpected error:", err)
				}
				fmt.Fprint(w, resp)
			}))
			defer s.Close()
			defer unexported.SetBaseURL(s.URL)() // baseURLをモックサーバのものに入れ替え

			cli := unexported.Client{HTTPClient: s.Client()}
			_, err := cli.Get(tc.n)
			switch {
			case err != nil && !tc.hasError:
				t.Error("unexpected error:", err)
			case err == nil && tc.hasError:
				t.Error("expected error has not occurred")
			}

			if !requested {
				t.Error("no request")
			}
		})
	}
}
