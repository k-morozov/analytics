package api

import (
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	host   string
	port   string
}

func NewClient(rawUrl string) (client ConnectClickHouse, err error) {
	u, err := url.Parse(rawUrl)
	if nil != err {
		return
	}

	client = &Client{
		client: &http.Client{},
		host:   u.Hostname(),
		port:   u.Port()}

	return
}

func (s *Client) Ping() (ok bool, err error) {
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

	resp, err := s.client.Do(request)
	if nil != err {
		return
	}

	if 200 == resp.StatusCode && "200 OK" == resp.Status {
		return true, nil
	}

	return
}

func (s *Client) CreateTables() (err error) {
	reqArgs := url.Values{}
	reqArgs.Add("query", "CREATE TABLE IF NOT EXISTS events("+
		"AppName TEXT,"+
		"AppVersion TEXT,"+
		"ClientId TEXT,"+
		"Action TEXT,"+
		"Category TEXT,"+
		"Label TEXT,"+
		"Value TEXT) ENGINE = MergeTree() ORDER BY AppName")

	reqUrl, err := url.Parse("http://" + s.host + ":" + s.port)
	if nil != err {
		return
	}
	//reqUrl.Path = "/query"
	reqUrl.RawQuery = reqArgs.Encode()
	//log.Println(reqUrl)

	request, _ := http.NewRequest("POST", reqUrl.String(), nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 AppleWebKit/531.21.10 (KHTML, like Gecko)")
	//request.Header.Set("Content-Type", "application/json; charset=utf-8; x-readonly=true")

	_, err = s.client.Do(request)
	if nil != err {
		return
	}
	//log.Println(resp)

	return nil
}

func (s *Client) AddMetrics(r CollectRequest) (err error) {
	reqArgs := url.Values{}
	reqArgs.Add("query", "INSERT INTO events(AppName, AppVersion, ClientId, Action, Category, Label, Value) "+
		" VALUES ("+
		"'"+r.AppName+"', "+
		"'"+r.AppVersion+"', "+
		"'"+r.ClientId+"', "+
		"'"+r.Action+"', "+
		"'"+r.Category+"', "+
		"'"+r.Label+"', "+
		"'"+r.Value+"' "+
		")")
	reqUrl, err := url.Parse("http://" + s.host + ":" + s.port)
	if nil != err {
		return
	}
	reqUrl.RawQuery = reqArgs.Encode()

	request, _ := http.NewRequest("POST", reqUrl.String(), nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 AppleWebKit/531.21.10 (KHTML, like Gecko)")
	//request.Header.Set("Content-Type", "application/json; charset=utf-8; x-readonly=true")

	resp, err := s.client.Do(request)
	if nil != err {
		return
	}
	log.Println(resp)

	return
}
