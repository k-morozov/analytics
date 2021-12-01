package main

import (
	"analytics/api"
	"log"
	"net/http"
	"time"
)

func main() {
	// @TBD args
	connection, _ := api.NewClient("http://13.49.159.232:8123")

	if status, err := connection.Ping(); !status {
		log.Fatalf("Ping failed: %v", err)
	}

	log.Println("ping: OK")

	err := connection.CreateTables()
	if nil != err {
		log.Fatalf("CreateTables: %v", err)
	}

	log.Println("tables: READY")

	//request := api.CollectRequest{
	//	AppName: "app",
	//	AppVersion: "0.1",
	//	ClientId: "0001",
	//	Action: "test",
	//	Category: "useless",
	//	Label: "goto",
	//	Value: "5",
	//}
	//
	//err = connection.AddMetrics(request)
	//if nil != err {
	//	log.Fatalf("CreateTables: %v", err)
	//}
	//log.Println("AddMetrics: Done")

	handler := &api.Handler{
		Con: connection,
	}

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
