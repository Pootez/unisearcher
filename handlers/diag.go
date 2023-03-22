package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"unisearcher/utils"
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

// DiagHandler handles requests to the diag endpoint
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Set content type
		http.Header.Add(w.Header(), "content-type", "application/json; charset=utf-8")

		// Get response from APIs
		resU, err := http.Get(utils.UniversitiesApi)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resU.Body.Close()

		resC, err := http.Get(utils.CountriesApi+"/v2/name/peru")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resC.Body.Close()

		// Create diag struct
		diag := utils.Diag{
			UniApi:       resU.StatusCode,
			CountriesApi: resC.StatusCode,
			Version:      utils.Version,
			Uptime:       uptime(),
		}

		// Write diag struct to response
		resData, err := json.MarshalIndent(diag,"","        ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(resData)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}