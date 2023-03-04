package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"unisearcher/utils"
	"io"
)

// startTime is the time when the server started
var startTime = time.Now()

// uptime returns the uptime of the server
func uptime() string {
	return time.Since(startTime).String()
}

// init sets the start time
func init() {
	startTime = time.Now()
}

// UniSearcherHandler handles requests to the UniSearcher service
func UniSearcherHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type
	http.Header.Add(w.Header(), "content-type", "text/html; charset=utf-8")

	// Handle request
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Welcome to Unisearcher service!<br>")
		fmt.Fprintf(w, "Available endpoints:<br>")
		fmt.Fprintf(w, "<a href=\"uniinfo\">uniinfo</a><br>")
		fmt.Fprintf(w, "<a href=\"neighbourunis\">neighbourunis</a><br>")
		fmt.Fprintf(w, "<a href=\"diag\">diag</a><br>")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// UniInfoHandler handles requests to the uniinfo endpoint
func UniInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Set content type
		http.Header.Add(w.Header(), "content-type", "application/json; charset=utf-8")

		// Get query
		l := len(strings.Split(utils.UniInfoPath, "/")) - 1
		query := strings.Replace(strings.Split(r.URL.Path, "/")[l], " ", "%20", -1)

		// Get response from API
		res, err := http.Get(utils.UniversitiesApi + "/search?name=" + query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return			
		}
		defer res.Body.Close()

		// Write response to response
		jsonData, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(jsonData)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// NeighbourHandler handles requests to the neighbourunis endpoint
func NeighbourHandler(w http.ResponseWriter, r *http.Request) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "This endpoint isn't implemented yet.")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// DiagHandler handles requests to the diag endpoint
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Set content type
		http.Header.Add(w.Header(), "content-type", "application/json; charset=utf-8")

		// Get response from APIs
		chanU := make(chan http.Response)
		go utils.GetRequest(utils.UniversitiesApi, chanU)
		respU := <-chanU

		chanC := make(chan http.Response)
		go utils.GetRequest(utils.CountriesApi, chanC)
		respC := <-chanC

		// Create diag struct
		diag := utils.Diag{
			UniApi:       respU.StatusCode,
			CountriesApi: respC.StatusCode,
			Version:      utils.Version,
			Uptime:       uptime(),
		}

		// Write diag struct to response
		jsonData, err := json.Marshal(diag)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(jsonData)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
