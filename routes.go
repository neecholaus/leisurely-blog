package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func getLanding(w http.ResponseWriter, r *http.Request) {
	html, _ := template.ParseFiles("./html/core-template.html", "./html/landing.html")
	_ = html.Execute(w, nil)
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	html, _ := template.ParseFiles("./html/core-template.html", "./html/about.html")
	_ = html.Execute(w, nil)
}

func getCreateDraft(w http.ResponseWriter, r *http.Request) {
	html, _ := template.ParseFiles("./html/core-template.html", "./html/create-draft.html")
	_ = html.Execute(w, nil)
}

type saveDraft_body struct {
	LastUsedTitle string `json:"lastUsedTitle"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}

func postSaveDraft(w http.ResponseWriter, r *http.Request) {
	var saveDraft_body saveDraft_body
	err := json.NewDecoder(r.Body).Decode(&saveDraft_body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error - unmarshal: %s", err.Error())))
		return
	}

	if saveDraft_body.Title == "" {
		saveDraft_body.Title = "Untitled"
	}
	if saveDraft_body.LastUsedTitle == "" {
		saveDraft_body.LastUsedTitle = "Untitled"
	}

	fmt.Printf("storing draft: %s\n", saveDraft_body.Title)

	if saveDraft_body.LastUsedTitle != saveDraft_body.Title {
		// todo - check if new title already exists.
		os.Rename(
			fmt.Sprintf("./drafts/%s", saveDraft_body.LastUsedTitle),
			fmt.Sprintf("./drafts/%s", saveDraft_body.Title),
		)
	}

	os.WriteFile(fmt.Sprintf("./drafts/%s", saveDraft_body.Title), []byte(saveDraft_body.Content), fs.FileMode(0777))

	response := struct {
		Title string `json:"title"`
	}{
		Title: saveDraft_body.Title,
	}
	encodedResponse, err := json.Marshal(&response)

	w.Header().Add("content-type", "application/json")
	w.Write(encodedResponse)
}

func getViewDraft(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	vars := mux.Vars(r)
	title, ok := vars["title"]
	if ok == false {
		w.Write([]byte("cannot"))
		return
	}

	file, err := os.ReadFile(fmt.Sprintf("./drafts/%s", title))
	if err != nil {
		w.Write([]byte("cannot"))
		return
	}

	w.Header().Add("content-type", "text/html")
	w.Write([]byte(title))
	w.Write([]byte("<br/>"))
	w.Write(file)
	return
}
