package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DeviceAuthorization struct {
	UserCode        string
	VerificationUri string
	ExpirationTime  int
}

type deviceAuthResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUri string `json:"verification_uri"`
	ExpirationTime  int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

func (resp deviceAuthResponse) toDeviceAuthorization() DeviceAuthorization {
	return DeviceAuthorization{
		UserCode:        resp.UserCode,
		VerificationUri: resp.VerificationUri,
		ExpirationTime:  resp.ExpirationTime,
	}
}

type DeviceAccess struct {
	Token string
}

type deviceAccessResponse struct {
	AccessToken string          `json:"access_token"`
	TokenType   string          `json:"token_type"`
	Scope       string          `json:"scope"`
	Error       accessErrorCode `json:"error"`
}

func (resp deviceAccessResponse) toDeviceAccess() DeviceAccess {
	return DeviceAccess{Token: resp.AccessToken}
}

type accessErrorCode string

const (
	authorizationPending accessErrorCode = "authorization_pending"
	slowDown             accessErrorCode = "slow_down"
	accessDenied         accessErrorCode = "access_denied"
	expiredToken         accessErrorCode = "expired_token"
	empty                accessErrorCode = ""
)

func (client *Client) DeviceFlow(ctx context.Context, errChan chan<- error, authChan chan<- DeviceAuthorization, accessChan chan<- DeviceAccess) {
	authResp, err := client.deviceAuthRequest(ctx)
	if err != nil {
		errChan <- err
		return
	}

	authChan <- authResp.toDeviceAuthorization()
	client.config.deviceCode = authResp.DeviceCode

	ttl := time.Duration(authResp.ExpirationTime) * time.Second
	interval := time.Duration(authResp.Interval) * time.Second

	ctx, cancel := context.WithTimeout(ctx, ttl)
	defer cancel()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
			return
		case <-ticker.C:
			acccessResp, err := client.deviceAccessRequest(ctx)
			if err != nil {
				errChan <- fmt.Errorf("oauth device polling error: %w", err)
				return
			}

			switch acccessResp.Error {
			case authorizationPending:
				ticker.Reset(interval)
			case slowDown:
				interval += 5 * time.Second
				ticker.Reset(interval)
			case accessDenied:
				errChan <- AccessDenied
				return
			case expiredToken:
				errChan <- ExpiredToken
				return
			case empty:
				accessChan <- acccessResp.toDeviceAccess()
				return
			}
		}
	}
}

func (client *Client) deviceAuthRequest(ctx context.Context) (*deviceAuthResponse, error) {
	body := url.Values{
		"client_id": {client.config.clientID},
		"scope":     client.config.scopes,
	}.Encode()

	url := client.config.deviceAuthURL
	if url == "" {
		return nil, fmt.Errorf("missing device authentification URL")
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, strings.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create device authentification request: %w", err)
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json"},
	}

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot do device authentification request: %w", err)
	}

	auth := &deviceAuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(auth)
	if err != nil {
		return nil, fmt.Errorf("error decode device authentification response: %w", err)
	}

	return auth, nil
}

func (client *Client) deviceAccessRequest(ctx context.Context) (*deviceAccessResponse, error) {
	body := url.Values{
		"client_id":   {client.config.clientID},
		"device_code": {client.config.deviceCode},
		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
	}.Encode()

	url := client.config.deviceAccessURL
	if url == "" {
		return nil, fmt.Errorf("missing device access URL")
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, strings.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create device access request: %w", err)
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json"},
	}

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot do device access request: %w", err)
	}

	access := &deviceAccessResponse{}
	err = json.NewDecoder(resp.Body).Decode(access)
	if err != nil {
		return nil, fmt.Errorf("cannot decode device access response: %w", err)
	}

	return access, nil
}
