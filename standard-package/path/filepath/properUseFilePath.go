package filepath

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cwd) // $HOME/go/src/github.com/mafuyuk/go-training/minna-no-go

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// httpリクエストは論理パスなのでpath
		if ok, err := path.Match("/data/*.html", r.URL.Path); err != nil || !ok {
			http.NotFound(w, r)
			return
		}

		// 以降は物理パスなのでpath/filepathを使う
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))
		// pathパッケージを使ってしまうと\がパスとしてみなされないため
		// http://localhost:8080/data/..\main.goなどで本来公開されていないファイルを
		// 見たりできるようになるので危険

		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}
