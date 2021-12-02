package services

import (
	"analytics/internal/interfaces"
	"analytics/internal/model"
	"analytics/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	// config
	// logger
	// middleware connection
	Con interfaces.ConnectClickHouse
}

func (h *Handler) Collect(response http.ResponseWriter, request *http.Request) {
	log.Println("Collect request")
	var args map[string]string
	args, err := utils.ParseQuery(request.URL.RawQuery)
	if nil != err {
		response.WriteHeader(400)
		log.Fatalf("error: %v", err)
	}

	r, err := Convert(args)
	log.Println(*r)

	err = h.Con.AddMetrics(*r)
	if err != nil {
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(response).Encode(args); nil != err {
		log.Fatalf("error: %v", err)
	}
	response.WriteHeader(200)
}

func Convert(args map[string]string) (r *model.CollectRequest, err error) {
	marsh, err := json.Marshal(args)
	//log.Printf("marsh = %v\n", string(marsh))
	err = json.Unmarshal(marsh, &r)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	return
}
