package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type JsonData struct {
	Key   string
	Value string
}

func ServerRun() {
	http.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if contentTypeCheck(r) {
				jsonData, err := jsonParser(r)
				if err != nil {
					log.Println(err)
					http.Error(w, "Invalid JSON", http.StatusBadRequest)
				} else {
					fmt.Fprintf(w, "Path of file: %+v", jsonData.Value)
				}
			} else {
				log.Printf("Invalid Content-Type, \"application/json\" is needed")
			}
		}

	})
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func jsonParser(r *http.Request) (JsonData, error) {
	jsonData := JsonData{}
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		log.Println("Invalid JSON")
		return jsonData, err
	}
	return jsonData, nil
}

func contentTypeCheck(r *http.Request) bool {
	contentType := r.Header.Get("content-type")
	return strings.Contains(contentType, "application/json")
}
