package main

import (
	"analytics/api"
	"analytics/client"
	"log"
	"net/http"
	"time"
)

func main() {
	if status := client.Ping(); !status {
		log.Fatalf("Ping failed")
	}

	log.Println("OK")

	_ = client.Send()

	handler := &api.Handler{}
	mux := http.NewServeMux()
	mux.HandleFunc("/collect", handler.Collect)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
