package controllers

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Functions for svgs
// get the svg from directory
func GetSvgAsString(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return []byte{}, err
	}

	return file, nil
}

var svgElements map[string]string

func getWhitelistedSVGElements() (map[string]string, error) {
	file, err := os.Open("assets\\PossibleSVG.json")

	if err != nil {
		return map[string]string{}, nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []byte

	for scanner.Scan() {
		lines = append(lines, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		return map[string]string{}, err
	}

	if err := json.Unmarshal(lines, &svgElements); err != nil {
		return map[string]string{}, err
	}

	return svgElements, nil
}

func getKeyValueSvg(file string) (string, bool) {
	getWhitelistedSVGElements()
	value, ok := svgElements[file]
	return value, ok
}

///////////////////////////////////////////////////////////////////////////////

func GetSvg(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	filename := parts[len(parts)-1]

	value, ok := getKeyValueSvg(filename)

	if ok {
		svgFile, err := GetSvgAsString("assets\\svgs\\" + value + ".svg")
		if err != nil {
			http.Error(w, "what", http.StatusInternalServerError)
		}

		msg := Message{
			Message: string(svgFile),
		}

		jsonSvgPayload, err := json.Marshal(msg)

		if err != nil {
			http.Error(w, "what", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(jsonSvgPayload)
	} else {
		http.Error(w, "no bonita im sowwy", http.StatusInternalServerError)
	}

}
