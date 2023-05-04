package main

import (
	"html/template"
	"net/http"
)

func getLanding(w http.ResponseWriter, r *http.Request) {
	html, _ := template.ParseFiles("./html/core-template.html", "./html/landing.html")
	_ = html.Execute(w, nil)
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	html, _ := template.ParseFiles("./html/core-template.html", "./html/about.html")
	_ = html.Execute(w, nil)
}
