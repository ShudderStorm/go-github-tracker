package graphql

import (
	"net/http"
)

type HttpAgent interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	endpoint  string
	httpAgent HttpAgent
}
