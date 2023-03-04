package handlers

import (
	"net/http"
	"fmt"
	"unisearcher/utils"
)

func DefaultHandler (w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "text/html; charset=utf-8")

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "This endpoint isn't implemented yet. Try <a href=\"" + utils.UniSearcherPath + "\">the unisearcher service</a> instead.")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}