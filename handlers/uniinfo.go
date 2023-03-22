package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"unisearcher/utils"
)

// UniinfoHandler handles requests to the uniinfo endpoint
func UniinfoHandler(w http.ResponseWriter, r *http.Request) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Set content type
		http.Header.Add(w.Header(), "content-type", "application/json; charset=utf-8")

		// Get query
		l := len(strings.Split(utils.UniinfoPath, "/")) - 1
		query := strings.Replace(strings.Split(r.URL.Path, "/")[l], " ", "%20", -1)

		// Get response from API
		res, err := http.Get(utils.UniversitiesApi + "/search?name=" + query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		// Read response
		jsonData, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Parse response
		countries := []interface{}{}
		err = json.Unmarshal(jsonData, &countries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Write response to response
		resData, err := json.MarshalIndent(countries, "", "        ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(resData)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
