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
	var args map[string]string
	args, err := ParseQuery(request.URL.RawQuery)
	if nil != err {
		response.WriteHeader(400)
		log.Fatalf("error: %v", err)
	}

	r, err := db.Convert(args)
	log.Println(*r)

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(response).Encode(args); nil != err {
		log.Fatalf("error: %v", err)
	}
	response.WriteHeader(200)
}
