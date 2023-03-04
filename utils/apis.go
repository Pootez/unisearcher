package utils

import (
	"log"
	"net/http"
)

func GetRequest(url string, c chan http.Response) {
	// Create a new HTTP client
	client := &http.Client{}
	defer client.CloseIdleConnections()	

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error in creating request:", err.Error())
		close(c)
		return
	}

	// Send the request and get the response
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error in sending request:", err.Error())
		close(c)
		return
	}

	// Send the response to the channel
	c <- *response

	// Close the channel
	close(c)
}