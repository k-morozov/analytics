package api

import (
	"log"
	"net/http"
	"net/url"
)

type Handler struct{}

func (h *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("call handler")

	if "/help" == request.URL.Path {
		log.Println("help handler")
		return
	}

	if "/collect" == request.URL.Path {
		collectHandler(response, request)
	}
}

func collectHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Collect request")
	u, err := url.Parse(request.URL.RawQuery)
	if nil != err {
		response.WriteHeader(400)
		log.Fatalf("err: %v", err)
	}
	args := u.Query()
	log.Println(args)

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(200)
}
