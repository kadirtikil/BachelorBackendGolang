package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Functions for svgs
// get the svg from directory

// struct to make svg object out of query return value:
type SvgObject struct {
	id      int
	title   string
	svgText string
}

func GetSvgFromDB(w http.ResponseWriter, r *http.Request) {
	// get the filename from the route
	parts := strings.Split(r.URL.Path, "/")
	svg := parts[len(parts)-1]

	fmt.Println(svg)

	// call it just in case it hasnt been called in any other controller but actually i dont need to since
	// every interaction with the backend start with fetching the backend anyway.
	// which makes the application vulnerable. if someone manages to mim the connection, and would call this function, knowing the database connection hasnt been set up
	// this person could make the server crash.... Need to add some serious error handling.
	initDatabase()

	// Create the object to save fetched data
	var svgObj SvgObject

	// The query to fetch the row identified by title, limitted to 1
	query := "SELECT * FROM bachelordatabase.svg WHERE title = ? LIMIT 1;"

	// exec the query check for err, safe data through scan and reference to svgObj above.
	err := db.QueryRow(query, svg).Scan(&svgObj.id, &svgObj.title, &svgObj.svgText)

	if err != nil {
		fmt.Println(err)
	}

	// Prepare return value
	msg := Message{
		Message: svgObj.svgText,
	}

	// jsonify the msg object to send it back
	jsonData, err2 := json.Marshal(msg)

	if err2 != nil {
		fmt.Println(err2)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(jsonData)

}
