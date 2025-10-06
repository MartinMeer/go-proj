package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {

	var jsonData map[string]interface{}

	http.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			contentType := r.Header.Get("content-type")
			if strings.Contains(contentType, "application/json") {
				//var jsonData map[string]interface{}
				err := json.NewDecoder(r.Body).Decode(&jsonData)
				if err != nil {
					log.Println("Invalid JSON")
					http.Error(w, "Invalid JSON", http.StatusBadRequest)
				} else {
					fmt.Fprintf(w, "Received JSON data: %s. Server is ready to read a file!", jsonData)

				}
			}
		}

	})

	log.Println("Starting server on http://localhost:8081")

	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(listener, nil))

}
