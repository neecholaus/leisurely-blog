package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getLanding)
	r.HandleFunc("/about", getAbout)

	_ = http.ListenAndServe(":8080", r)
}
