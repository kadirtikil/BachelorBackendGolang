package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func UpdateMarkdownTxt(pathfile string) (string, error) {
	fmt.Println(pathfile)

	return "", nil
}

func fetchMarkdownAfterUpdate(filename string) (string, error) {

	return "", nil
}

// Check Creds to edit and manip the txt for the markdown
func CheckCreds(name string, password string) bool {
	if name == "this" && password == "that" {
		return true
	} else {
		return false
	}
}

// handler func
func UpdateMarkdownData(w http.ResponseWriter, r *http.Request) {
	// get the filename from the route
	parts := strings.Split(r.URL.Path, "/")
	filename := parts[len(parts)-1]

	// The msg object from the struct to safe the new markdown from the json in request body in.
	var msg Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		http.Error(w, "upsie daisy", http.StatusInternalServerError)
	}

	// just printing it to check.
	// fmt.Println(msg)

	// check if creds are ok
	if CheckCreds("this", "that") {
		// Cred are ok so go and update the old markdown txt.
		newMarkdown, err := UpdateMarkdownTxt("" + filename + ".txt")

		if err != nil {
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		}

		// fetch the new fresh markdown and send it back. Dont want to do it clientside cause then some sync stuff might be going no bueno.

		fmt.Fprintf(w, newMarkdown)
	} else {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
	}

}
