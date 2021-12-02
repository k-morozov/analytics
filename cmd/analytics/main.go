package main

import (
	"analytics/internal/services"
	"log"
)

func main() {
	// @TBD args
	client, _ := services.NewClient("http://13.49.159.232:8123")

	if status, err := client.Ping(); !status {
		log.Fatalf("Ping failed: %v", err)
	}

	log.Println("ping: OK")

	err := client.CreateTables()
	if nil != err {
		log.Fatalf("CreateTables: %v", err)
	}

	server := services.NewEventService(client)

	log.Println("tables: READY")

	if err := server.EventReceiver.ListenAndServe(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
