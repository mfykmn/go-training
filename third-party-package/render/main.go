package main

import (
	"net/http"

	"github.com/unrolled/render"
)

// $ curl http://localhost:8080/
// {"reason":"success"}
var rendering = render.New(render.Options{})

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type response struct {
	Reason string `json:"reason"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	rendering.JSON(w, http.StatusNotFound, &response{Reason: "success"})
}
