package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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

	// this is where the chifre check will take place...
	if moe.Chifre1 == "Hello" && moe.Chifre2 == "World" {
		fmt.Println("Password is correct!!!")
		// all the data needet is here. now its time to update the entry in the db
		// query the data from the table
		query := "UPDATE markdown SET markdown = ? WHERE title = ?"
		res, err := db.Exec(query, moe.MarkdownText, moe.Headline)

		if err != nil {
			fmt.Print(err)
		}

		fmt.Println("Row has been updated", res)
	} else {
		http.Error(w, "sowwy", http.StatusInternalServerError)
	}

	// Respond with the message
	//fmt.Println(jsonData)
	// Optionally, you can send a response back to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message received successfully"))
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Define struct for request
type CheckAuthObject struct {
	Chiffre1 string `json:"chiffre1"`
	Chiffre2 string `json:"chiffre2"`
}

// Define the struct for the response
type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Println(string(content))

	if err != nil {
		fmt.Println("something wrong trying to read the request")
		http.Error(w, "No Auth possible at this time. Contact DEV if issue persists.", http.StatusInternalServerError)
		return
	}

	// at this point, content is byte data from the request body, so creds reach the backend.
	var cao CheckAuthObject
	err2 := json.Unmarshal(content, &cao)

	if err2 != nil {
		fmt.Println("something wrong trying to unmarshal the JSON at CheckAuth func")
		http.Error(w, "No Auth possible at this time. Contact DEV if issue persists.", http.StatusInternalServerError)
		return
	}

	// chiffres are reaching the backend now. time to check them.
	if cao.Chiffre1 == "Hello" && cao.Chiffre2 == "World" {
		fmt.Println("all good in the hood")
		json.NewEncoder(w).Encode(AuthResponse{Success: true, Message: "All good"})
	} else {
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "no bueno"})
	}

	fmt.Println(cao)
}
