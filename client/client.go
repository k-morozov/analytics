package client

import (
	"log"
	"net/http"
	"net/url"
)

func Ping() bool {
	client := &http.Client{}

	reqArgs := url.Values{}
	//reqArgs.Add("ping", "")
	reqUrl, err := url.Parse("http://13.49.159.232:8123")
	if nil != err {
		log.Fatalf("parse: %v", err)
		return false
	}
	reqUrl.Path = "/ping"
	reqUrl.RawQuery = reqArgs.Encode()

	request, _ := http.NewRequest("GET", reqUrl.String(), nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 AppleWebKit/531.21.10 (KHTML, like Gecko)")

	resp, err := client.Do(request)
	if nil != err {
		log.Fatalf("do: %v", err)
		return false
	}

	log.Println(resp)
	return true
}
