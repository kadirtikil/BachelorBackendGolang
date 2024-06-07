package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/******************************************************************************************************************************************************/
//////////////////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

// Struct for response format
type Message struct {
	Message string `json:"message"`
}

// Struct for fetching markdown from db
type Markdown struct {
	id       int64
	title    string
	markdown string
}

func GetMarkdownFromDB(w http.ResponseWriter, r *http.Request) {
	// get the filename from the route
	parts := strings.Split(r.URL.Path, "/")
	record := parts[len(parts)-1]

	initDatabase()
	// return value
	var mrkdn Markdown

	// query the data from the table
	query := "SELECT * FROM markdown WHERE title = ? LIMIT 1"
	err := db.QueryRow(query, record).Scan(&mrkdn.id, &mrkdn.title, &mrkdn.markdown)

	if err != nil {
		fmt.Println("Error")
	}

	msg := Message{
		Message: mrkdn.markdown,
	}

	jsonData, err2 := json.Marshal(msg)

	if err2 != nil {
		http.Error(w, "some error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(jsonData)
}
