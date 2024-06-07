package controllers

import (
	"fmt"
	"io"
	"net/http"
)

func UpdateMarkdown(w http.ResponseWriter, r *http.Request) {

	content, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("error in reequst")
	}

	defer r.Body.Close()

	msg := Message{
		Message: string(content),
	}

	// string(content) now contains the message as string. but why not as json? time to decode som stuff.

	fmt.Println(msg)
}
