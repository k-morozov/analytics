package main

import (
	"analytics/api"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Hello world!")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        &api.Handler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
