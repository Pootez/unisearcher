package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"unisearcher/utils"
	"strconv"
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
		resData, err := json.MarshalIndent(countries,"","        ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(resData)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// NeighbourHandler handles requests to the neighbourunis endpoint
func NeighbourHandler(w http.ResponseWriter, r *http.Request) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Set content type
		http.Header.Add(w.Header(), "content-type", "application/json; charset=utf-8")

		// Get query
		l := len(strings.Split(utils.UniInfoPath, "/")) - 1
		queryArr := strings.Split(strings.Replace(r.URL.Path, " ", "%20", -1), "/")
		queryParams := r.URL.Query()
		limitQuery, limitBool := queryParams["limit"]
		var limit int
		var err error
		if limitBool {
			limit, err = strconv.Atoi(limitQuery[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Check if query is valid
		if len(strings.Split(r.URL.Path, "/")) <= l+1 {
			http.Error(w, "Missing url parameters", http.StatusBadRequest)
			return
		}

		// Get response from API
		resCountry, err := http.Get(utils.CountriesApi + "/v3.1/name/" + queryArr[l] + "?fullText=true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resCountry.Body.Close()

		// Read response
		jsonCountry, err := io.ReadAll(resCountry.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Unmarshal json
		var country []utils.Country
		err = json.Unmarshal(jsonCountry, &country)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Check if country is valid
		if len(country) == 0 {
			http.Error(w, "Country not found", http.StatusNotFound)
			return
		}

		// Check if country has neighbours
		if len(country[0].Borders) == 0 {
			http.Error(w, "Country has no neighbours", http.StatusNotFound)
			return
		}

		// Get neighbours from API
		neighbours := []utils.Country{}

		for _, border := range country[0].Borders {
			resNeighbour, err := http.Get(utils.CountriesApi + "/v3.1/alpha/" + border)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer resNeighbour.Body.Close()

			// Read response
			jsonNeighbour, err := io.ReadAll(resNeighbour.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			// Unmarshal json
			var neighbour []utils.Country
			err = json.Unmarshal(jsonNeighbour, &neighbour)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				break
			}

			// Check if neighbour is valid
			if len(neighbour) != 0 {
				neighbours = append(neighbours, neighbour[0])
			}
		}

		// Get universities from neighbours
		universities := []interface{}{}

		for _, neighbour := range neighbours {
			resUni, err := http.Get(utils.UniversitiesApi + "/search?name=" + queryArr[l+1] + "&country=" + neighbour.Name.Common)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer resUni.Body.Close()

			// Read response
			jsonUni, err := io.ReadAll(resUni.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			// Unmarshal universities
			var unis []interface{}
			err = json.Unmarshal(jsonUni, &unis)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				break
			}

			// Limit universities
			if len(unis) != 0 {
				if limitBool && len(unis) > limit {
					unis = unis[:limit]
				}
				universities = append(universities, unis...)
			}
		}

		// Write response to response
		resData, err := json.MarshalIndent(universities,"","        ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(resData)
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
