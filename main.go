package main

import (
	"log"
	"net/http"
)

func collectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("called handler")
}

func main() {
	log.Println("Hello world!")

	http.HandleFunc("/collect", collectHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %v", err)
	}
}
