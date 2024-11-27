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

type Handler struct {
	doer   Doer
	config config
}

type Option func(*Handler)

func New(doer Doer, opts ...Option) *Handler {
	handler := &Handler{
		doer: doer,
	}

	for _, opt := range opts {
		opt(handler)
	}

	return handler
}

func WithClientID(clientID string) Option {
	return func(h *Handler) {
		h.config.clientID = clientID
	}
}

func WithDeviceAuthURL(url string) Option {
	return func(h *Handler) {
		h.config.deviceAuthURL = url
	}
}

func WithDeviceAccessURL(url string) Option {
	return func(h *Handler) {
		h.config.deviceAccessURL = url
	}
}

func WithScopes(scopes ...string) Option {
	return func(h *Handler) {
		h.config.scopes = scopes
	}
}
