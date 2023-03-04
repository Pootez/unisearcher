package handlers

import (
	"net/http"
	"fmt"
	"unisearcher/utils"
)

// DefaultHandler handles requests to the default path
func DefaultHandler (w http.ResponseWriter, r *http.Request) {
	// Set content type
	http.Header.Add(w.Header(), "content-type", "text/html; charset=utf-8")

	// Handle request
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "This endpoint isn't implemented yet. Try <a href=\"" + utils.UniSearcherPath + "\">the unisearcher service</a> instead.")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}