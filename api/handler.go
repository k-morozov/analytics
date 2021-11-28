package api

import (
	"analytics/db"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	// config
	// logger
	// db connection
}

func (h *Handler) Collect(response http.ResponseWriter, request *http.Request) {
	log.Println("Collect request")
	args, err := ParseQuery(request.URL.RawQuery)
	if nil != err {
		response.WriteHeader(400)
		log.Fatalf("error: %v", err)
	}
	log.Printf("args = %v\n", args)

	marsh, err := json.Marshal(args)
	log.Printf("marsh = %v\n", string(marsh))
	r := db.CollectRequest{}
	err = json.Unmarshal(marsh, &r)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	log.Printf("CollectRequest = %v\n", r)

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(response).Encode(args); nil != err {
		log.Fatalf("error: %v", err)
	}
	response.WriteHeader(200)
}
