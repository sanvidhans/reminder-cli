package client

import "net/http"

type HTTPClient struct {
	client *http.Client
	BackendURL string
}

func NewHTTPClient(url string) HTTPClient {
	return HTTPClient{
		client:     &http.Client{},
		BackendURL: url,
	}
}