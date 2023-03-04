package handlers

import (
	"net/http"
	"fmt"
)

// UniSearcherHandler handles requests to the UniSearcher service
func UniSearcherHandler (w http.ResponseWriter, r *http.Request) {
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