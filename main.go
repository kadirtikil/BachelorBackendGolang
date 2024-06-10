package main

import (
	"BachelorBackendGolang/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/svg/{svg}", controllers.GetSvgFromDB)

	mux.HandleFunc("/editmarkdown", controllers.UpdateMarkdown)

	mux.HandleFunc("/fetchmarkdown/{record}", controllers.GetMarkdownFromDB)

	mux.HandleFunc("/checkAuth", controllers.CheckAuth)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
