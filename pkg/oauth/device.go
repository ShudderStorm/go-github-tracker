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

type deviceAuthResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUri string `json:"verification_uri"`
	TTL             int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

type deviceAccessResponse struct {
	АccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Error       string `json:"error"`
}

type UserCodeRetriever func(string)

func (handler *Handler) DeviceFlow(ctx context.Context, retriever UserCodeRetriever) error {
	auth, err := handler.requestDeviceAuth(ctx)
	if err != nil {
		return fmt.Errorf("oauth device authentification error: %w", err)
	}

	retriever(auth.UserCode)
	handler.config.deviceCode = auth.DeviceCode

	ttl := time.Duration(auth.TTL) * time.Second
	interval := time.Duration(auth.Interval) * time.Second

	ctx, cancel := context.WithTimeout(ctx, ttl)
	defer cancel()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			auth, err := handler.requestDeviceAccess(ctx)
			if err != nil {
				return fmt.Errorf("oauth device access error: %w", err)
			}

			if auth.Error == "" {
				return nil
			} else if auth.Error != "authorization_pending" {
				return fmt.Errorf("oauth device access returns error code: %v", auth.Error)
			}
		}
	}
}

func (handler *Handler) requestDeviceAuth(ctx context.Context) (*deviceAuthResponse, error) {
	body := url.Values{
		"client_id": {handler.config.clientID},
		"scope":     handler.config.scopes,
	}.Encode()

	url := handler.config.deviceAuthURL
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

	resp, err := handler.doer.Do(req)
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

func (handler *Handler) requestDeviceAccess(ctx context.Context) (*deviceAccessResponse, error) {
	body := url.Values{
		"client_id":   {handler.config.clientID},
		"device_code": {handler.config.deviceCode},
		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
	}.Encode()

	url := handler.config.deviceAccessURL
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

	resp, err := handler.doer.Do(req)
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
