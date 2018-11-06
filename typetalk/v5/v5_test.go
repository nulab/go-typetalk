package v5

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const fixturesPath = "../../testdata/v5/"

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	parsedURL, _ := url.Parse(server.URL)
	client.client.BaseURL = parsedURL
}

func teardown() {
	server.Close()
}
