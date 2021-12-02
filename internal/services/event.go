package services

import (
	"analytics/internal/interfaces"
	"analytics/internal/model"
	"analytics/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type EventService struct {
	// config
	// logger
	// middleware connection
	EventReceiver *http.Server
	client        interfaces.ConnectClickHouse
}

func NewEventService(client interfaces.ConnectClickHouse) *EventService {

	service := &EventService{
		client: client,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/collect", service.collect)

	service.EventReceiver = &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return service
}

func (h *EventService) collect(response http.ResponseWriter, request *http.Request) {
	log.Println("Collect request")
	var args map[string]string
	args, err := utils.ParseQuery(request.URL.RawQuery)
	if nil != err {
		response.WriteHeader(400)
		log.Fatalf("error: %v", err)
	}

	event, err := convert(args)

	err = h.client.AddMetrics(*event)
	if err != nil {
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(response).Encode(args); nil != err {
		log.Fatalf("error: %v", err)
	}
	response.WriteHeader(200)
}

func convert(args map[string]string) (r *model.Event, err error) {
	marsh, err := json.Marshal(args)
	//log.Printf("marsh = %v\n", string(marsh))
	err = json.Unmarshal(marsh, &r)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	return
}
