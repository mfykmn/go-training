package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	// Root への処理
	r.HandleFunc("/", handler).Methods("GET")

	// アドミン周りの処理
	authBase := mux.NewRouter()
	r.PathPrefix("/admin/").Handler(negroni.New(
		negroni.HandlerFunc(MyMiddleware),
		negroni.Wrap(authBase)))
	auth := authBase.PathPrefix("/admin").Subrouter()
	auth.Path("/").HandlerFunc(adminHandler)

	// http://localhost:3000/admin/ リクエストを送るとdo middleware ⇒ called adminHandler という出力を得る

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3000")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called handler")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called adminHandler")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("do middleware")
	next(rw, r)
}
