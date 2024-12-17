package oauth

import "net/http"

type config struct {
	clientID        string
	scopes          []string
	deviceAuthURL   string
	deviceAccessURL string
	deviceCode      string
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	client Doer
	config config
}

type Option func(*Client)

func New(doer Doer, opts ...Option) *Client {
	client := &Client{
		client: doer,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithId(id string) Option {
	return func(c *Client) {
		c.config.clientID = id
	}
}

func WithDeviceFlow(authUrl string, pollUrl string) Option {
	return func(c *Client) {
		c.config.deviceAuthURL = authUrl
		c.config.deviceAccessURL = pollUrl
	}
}

func WithScopes(scopes ...string) Option {
	return func(c *Client) {
		c.config.scopes = scopes
	}
}
