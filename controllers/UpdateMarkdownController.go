package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

func UpdateMarkdownTxt(pathfile string) (string, error) {
	fmt.Println(pathfile)

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

	// check if creds are ok
	if CheckCreds("this", "that") {
		// fetch the new fresh markdown and send it back. Dont want to do it clientside cause then some sync stuff might be going no bueno.
		newMarkdown, err := UpdateMarkdownTxt("" + filename + ".txt")

		if err != nil {
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		}

		fmt.Fprintf(w, newMarkdown)
	} else {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
	}

}
