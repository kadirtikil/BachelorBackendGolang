package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MarkdownEditObject struct {
	// remember ALWAYS BEGIN WITH UPPERCASE!!!!!!
	Headline     string `json:"headline"`
	MarkdownText string `json:"markdown"`
	Chifre1      string `json:"chifre1"`
	Chifre2      string `json:"chifre2"`
}

func UpdateMarkdown(w http.ResponseWriter, r *http.Request) {
	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Set CORS headers for the actual request
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Process the request
	content, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error in request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var moe MarkdownEditObject

	err2 := json.Unmarshal((content), &moe)

	if err2 != nil {
		fmt.Println(err2)
	}

	// OOOOKAAAY now i got the data and can query the DB!
	// fmt.Println(moe.Chifre1)

	// all the data needet is here. now its time to update the entry in the db
	// query the data from the table
	query := "UPDATE markdown SET markdown = ? WHERE title = ?"

	// this is where the chifre check will take place...
	if true {
		res, err := db.Exec(query, moe.MarkdownText, moe.Headline)

		if err != nil {
			fmt.Print(err)
		}

		fmt.Println(res)
	}

	// Respond with the message
	//fmt.Println(jsonData)
	// Optionally, you can send a response back to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message received successfully"))
}
