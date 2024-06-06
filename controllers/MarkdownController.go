package controllers

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

/******************************************************************************************************************************************************/
//////////////////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS
// Gets the markdown txt
func GetTxtAsString(filepath string) (string, error) {
	// open the file in question
	file, err := os.Open(filepath)
	// check if error
	if err != nil {
		return "", err
	}
	//close when possible
	defer file.Close()

	// scan line and create buffer to safe lines
	scanner := bufio.NewScanner(file)
	var lines []string

	// go line by line append to buffer
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check for error in scanner
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// return the string
	return strings.Join(lines, "\n"), nil
}

// to safe the data from json
var jsonDataWhitelisted map[string]string

// read the json containing fetchable elements. (whitelisted)
func getWhitelistedElements() (map[string]string, error) {
	// open the json here
	file, err := os.Open("assets\\PossibleFiles.json")
	// check if json might not exist
	if err != nil {
		return map[string]string{}, err
	}
	// close whenever opened
	defer file.Close()

	// init line by line scan params
	scanner := bufio.NewScanner(file)
	var lines []byte

	for scanner.Scan() {
		lines = append(lines, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		return map[string]string{}, err
	}

	// checked if worked. it works. nice (before it was byte. needs to be byte to unmarshal as json)
	//fmt.Println(strings.Join(lines, "\n"))

	// jsonify the string. because it is a matrix.
	if err := json.Unmarshal(lines, &jsonDataWhitelisted); err != nil {
		return map[string]string{}, err
	}

	// This works now too. its a map now. so it can be fetched, and cheked at get requests. gg wp
	//fmt.Println(jsonDataWhitelisted)

	return jsonDataWhitelisted, nil
}

// Check if key is in map
func getKeyValue(key string) (string, bool) {
	// to init jsonDataWhitelisted, this func needs to be called. might not be the most efficient solution but its ok since the json is very small.
	// is it wasnt, then wouldve read peristently with some watchdog for the json, such that the var gets init if changes occur.
	getWhitelistedElements()
	value, ok := jsonDataWhitelisted[key]
	return value, ok
}

/******************************************************************************************************************************************************/
//////////////////////////////////////////////////////////////////////////////////////////

// Struct for response format

type Message struct {
	Message string `json:"message"`
}

func GetMarkdown(w http.ResponseWriter, r *http.Request) {
	// get the filename from the route
	parts := strings.Split(r.URL.Path, "/")
	filename := parts[len(parts)-1]

	// Check if file is whitelisted and get value of key
	value, ok := getKeyValue(filename)

	if ok {
		// if ok then continue
		content, err2 := GetTxtAsString("assets\\text\\" + value + ".txt")

		if err2 != nil {
			http.Error(w, "Could'nt read file!", http.StatusInternalServerError)
		}
		// trying to return it as json here
		msg := Message{
			Message: content,
		}

		jsonResponseData, err := json.Marshal(msg)

		if err != nil {
			http.Error(w, "here", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Write(jsonResponseData)

	} else {
		http.Error(w, "Could'nt read file!", http.StatusInternalServerError)
	}

}
