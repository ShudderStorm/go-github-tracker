package oauth

import (
	"net/http"
	"time"
)

type config struct {
	id     string
	secret string

	authURI  string
	tokenURI string

	redirectURI string
	scopes      []string

	withState bool
	validator StateValidator
	stateTTL  time.Duration
}

type Client struct {
	client *http.Client
	config config
}

type Option func(*Client)

func New(authUri string, tokenUri string, id string, secret string, opts ...Option) *Client {
	client := &Client{
		config: config{
			authURI:  authUri,
			tokenURI: tokenUri,
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
		c.config.redirectURI = uri
	}
}

func WithScopes(scopes ...string) Option {
	return func(c *Client) {
		c.config.scopes = scopes
	}
}

func WithState(validator StateValidator) Option {
	return func(c *Client) {
		c.config.withState = true
		c.config.validator = validator
		c.config.stateTTL = DefaultStateTTL
	}
}

func WithStateTTL(ttl time.Duration) Option {
	return func(c *Client) {
		if !c.config.withState {
			return
		}

		c.config.stateTTL = ttl
	}
}
