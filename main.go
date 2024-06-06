package main

import (
	"BachelorBackendGolang/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/markdown/{filename}", controllers.GetMarkdown)

	mux.HandleFunc("/svg/{filename}", controllers.GetSvg)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
