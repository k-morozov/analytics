package client

import (
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	host string
	port string
}

func NewClient(rawUrl string) (client ConnectClickHouse, err error) {
	u, err := url.Parse(rawUrl)
	if nil != err {
		return
	}

	return &Client{host: u.Hostname(), port: u.Port()}, nil
}

func (s *Client) Ping() (ok bool, err error) {
	client := &http.Client{}

	reqUrl, err := url.Parse("http://" + s.host + ":" + s.port)
	if nil != err {
		return
	}
	reqUrl.Path = "/ping"

	reqArgs := url.Values{}
	reqUrl.RawQuery = reqArgs.Encode()

	request, err := http.NewRequest("GET", reqUrl.String(), nil)
	if nil != err {
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 AppleWebKit/531.21.10 (KHTML, like Gecko)")

	resp, err := client.Do(request)
	if nil != err {
		return
	}

	if 200 == resp.StatusCode && "200 OK" == resp.Status {
		return true, nil
	}

	return
}

func (s *Client) Send() bool {
	client := &http.Client{}

	reqArgs := url.Values{}
	reqArgs.Add("query", "CREATE TABLE IF NOT EXISTS test(title TEXT) ENGINE = Memory")

	reqUrl, err := url.Parse("http://" + s.host + ":" + s.port)
	if nil != err {
		log.Fatalf("parse: %v", err)
		return false
	}
	//reqUrl.Path = "/query"
	reqUrl.RawQuery = reqArgs.Encode()
	//log.Println(reqUrl)

	request, _ := http.NewRequest("POST", reqUrl.String(), nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 AppleWebKit/531.21.10 (KHTML, like Gecko)")
	//request.Header.Set("Content-Type", "application/json; charset=utf-8; x-readonly=true")

	resp, err := client.Do(request)
	if nil != err {
		log.Fatalf("do: %v", err)
		return false
	}

	log.Println(resp)
	return true
}
