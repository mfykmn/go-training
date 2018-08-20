package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

// $ curl http://localhost:8080/ -XPOST -d '{"name":"asd","description":"aaa"}'
// {"reason":"success"}

// $ curl http://localhost:8080/ -XPOST -d '{"name":"asd","description":""}'
// {"reason":"failed validation"}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type postRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"required,min=1,max=20"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var data postRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"reason":"failed json decode"}`))
		return
	}

	// バリデーション
	if err := validate.Struct(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"reason":"failed validation"}`))
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"reason":"success"}`))
}
