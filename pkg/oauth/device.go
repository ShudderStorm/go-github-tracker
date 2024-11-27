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

		}
	}
}

func (handler *Handler) requestDeviceAuth(ctx context.Context) (*deviceAuthResponse, error) {
	body := url.Values{
		"client_id": {handler.config.clientID},
		"scope":     {strings.Join(handler.config.scopes, "")},
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

func (handler *Handler) requestDeviceAccess(ctx context.Context) (deviceAccessResponse, error) {
	return deviceAccessResponse{}, nil
}
