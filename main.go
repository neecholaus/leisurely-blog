package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/css/"))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fs))

	r.HandleFunc("/", getLanding)
	r.HandleFunc("/about", getAbout)
	r.HandleFunc("/create-draft", getCreateDraft)
	r.HandleFunc("/save-draft", postSaveDraft)

	_ = http.ListenAndServe(":8080", r)
}
