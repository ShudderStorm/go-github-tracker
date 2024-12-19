package github

import (
	"context"
	"net/http"

	"github.com/ShudderStorm/go-github-tracker/pkg/oauth"
)

const (
	AuthURL   string = "https://github.com/login/device/code"
	AccessURL string = "https://github.com/login/oauth/access_token"
)

type ClientDevice struct {
	oauth *oauth.Client

	Authorization <-chan oauth.DeviceAuthorization
	Access        <-chan oauth.DeviceAccess
	Error         <-chan error

	authChan   chan oauth.DeviceAuthorization
	accessChan chan oauth.DeviceAccess
	errorChan  chan error
}

func (device *ClientDevice) StartAuthorization(ctx context.Context) {
	device.oauth.DeviceFlow(ctx, device.errorChan, device.authChan, device.accessChan)
}

func NewClientDevice(id string, scopes ...Scope) *ClientDevice {
	device := &ClientDevice{
		authChan:   make(chan oauth.DeviceAuthorization),
		accessChan: make(chan oauth.DeviceAccess),
		errorChan:  make(chan error),
	}

	device.Authorization = device.authChan
	device.Access = device.accessChan
	device.Error = device.errorChan

	oauthScopes := make([]string, len(scopes))
	for i, s := range scopes {
		oauthScopes[i] = string(s)
	}

	device.oauth = oauth.New(
		http.DefaultClient,
		oauth.WithDeviceFlow(AuthURL, AccessURL),
		oauth.WithId(id),
		oauth.WithScopes(oauthScopes...),
	)

	return device
}

type Scope string

const (
	Repo string = "repo"
)
