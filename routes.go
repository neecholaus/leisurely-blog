package main

import "net/http"

func getLanding(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("landing page"))
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("about page"))
}
