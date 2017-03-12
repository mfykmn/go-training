package main

import (
	"github.com/golang/go/src/os"
	"io"
	"net/http"
	"path"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", upload)

}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stream, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		p := filepath.Join("files", filepath.Base(header.Filename)) // filepathじゃないと\が送られても素通しでbad
		println(p)
		f, err := os.Create(p)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		io.Copy(f, stream)
		http.Redirect(w, r, path.Join("/file", p), 301)

	} else {
		http.Redirect(w, r, "/", 301)
	}
}
