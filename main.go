package main

import (
	"BachelorBackendGolang/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/svg/{filename}", controllers.GetSvg)

	mux.HandleFunc("/editmarkdown", controllers.UpdateMarkdownData)

	mux.HandleFunc("/fetchmarkdown/{record}", controllers.GetMarkdownFromDB)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
