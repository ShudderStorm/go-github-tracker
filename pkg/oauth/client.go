package oauth

import "net/http"

type config struct {
	id     string
	secret string

	authEndpoint  string
	tokenEndpoint string

	redirectUri string
	scopes      []string
}

type Client struct {
	client *http.Client
	config config
}

type Option func(*Client)

func New(authEndpoint string, tokenEndpoint string, opts ...Option) *Client {
	client := &Client{
		config: config{
			authEndpoint:  authEndpoint,
			tokenEndpoint: tokenEndpoint,
		},
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithId(id string) Option {
	return func(c *Client) {
		c.config.id = id
	}
}

func WithSecret(secret string) Option {
	return func(c *Client) {
		c.config.secret = secret
	}
}

func WithRedirectUri(uri string) Option {
	return func(c *Client) {
		c.config.redirectUri = uri
	}
}

func WithScopes(scopes ...string) Option {
	return func(c *Client) {
		c.config.scopes = scopes
	}
}
