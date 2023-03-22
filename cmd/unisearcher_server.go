package main

import (
	"log"
	"net/http"
	"os"
	"unisearcher/handlers"
	"unisearcher/utils"
)

func main() {

	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = utils.DefaultPort
	}

	// Define handlers
	http.HandleFunc(utils.DefaultPath, handlers.DefaultHandler)
	http.HandleFunc(utils.UniSearcherPath, handlers.UniSearcherDefaultHandler)
	http.HandleFunc(utils.UniinfoPath, handlers.UniinfoHandler)
	http.HandleFunc(utils.NeighbourPath, handlers.NeighbourHandler)
	http.HandleFunc(utils.DiagPath, handlers.DiagHandler)

	// Start server
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
