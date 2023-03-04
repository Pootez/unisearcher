package main

import (
	"os"
	"net/http"
	"log"
	"unisearcher/utils"
	"unisearcher/handlers"
)

func main() {

	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = utils.DefaultPort
	}

	http.HandleFunc(utils.DefaultPath, handlers.DefaultHandler)
	http.HandleFunc(utils.UniSearcherPath, handlers.UniSearcherHandler)
	
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}