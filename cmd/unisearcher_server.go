package main

import (
	"os"
	"log"
	"net/http"
	"unisearcher/utils"
)

func main() {

	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = utils.DefaultPort
	}
	
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}