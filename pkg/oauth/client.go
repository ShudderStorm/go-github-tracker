package oauth

import "net/http"

type config struct {
	id     string
	secret string

	authUri  string
	tokenUri string

	redirectUri string
	scopes      []string
}

type Client struct {
	client *http.Client
	config config
}

type Option func(*Client)

func New(authUri string, tokenUri string, id string, secret string, opts ...Option) *Client {
	client := &Client{
		config: config{
			authUri:  authUri,
			tokenUri: tokenUri,
			id:       id,
			secret:   secret,
		},
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
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
