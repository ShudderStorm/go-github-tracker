package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Access struct {
	Token  string
	Type   string
	Scopes []string
}

type accessResponse struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
	Scope string `json:"scope"`
}

func (r *accessResponse) toAccess() Access {
	return Access{
		Token:  r.Token,
		Type:   r.Type,
		Scopes: strings.Split(r.Scope, ","),
	}
}

func (c *Client) GetAuthorizationUrl() (string, error) {
	if c.config.id == "" {
		return "", fmt.Errorf("%w:%s", ConfigError, "'id' parameter cannot be empty")
	}

	_, err := url.ParseRequestURI(c.config.authUri)
	if err != nil {
		return "", fmt.Errorf("%w:%s:%w", ConfigError, "authorization URI validation error", err)
	}

	params := url.Values{
		"client_id":     {c.config.id},
		"response_type": {"token"},
		"scopes":        c.config.scopes,
	}

	if c.config.redirectUri != "" {
		_, err = url.ParseRequestURI(c.config.redirectUri)
		if err != nil {
			return "", fmt.Errorf("%w:%s:%w", ConfigError, "redirect URI validation error", err)
		}

		params.Add("redirect_uri", c.config.redirectUri)
	}

	return fmt.Sprintf("%s?%s", c.config.authUri, params.Encode()), nil
}

func (c *Client) Exchange(ctx context.Context, code string) (Access, error) {
	var (
		access Access
		err    error
	)

	params := url.Values{
		"client_id":     {c.config.id},
		"client_secret": {c.config.secret},
		"code":          {code},
	}

	if c.config.redirectUri != "" {
		_, err = url.ParseRequestURI(c.config.redirectUri)
		if err != nil {
			return access, fmt.Errorf("%w:%s:%w", ConfigError, "redirect URI validation error", err)
		}

		params.Add("redirect_uri", c.config.redirectUri)
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, c.config.tokenUri, strings.NewReader(params.Encode()),
	)

	if err != nil {
		return access, fmt.Errorf("%w:%w", RequestError, err)
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json"},
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return access, fmt.Errorf("%w:%w", RequestError, err)
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return access, fmt.Errorf("%w: response status code is %v", ClientError, resp.StatusCode)
	}

	if resp.StatusCode >= 500 {
		return access, fmt.Errorf("%w: response status code is %v", ServerError, resp.StatusCode)
	}

	accessResp := &accessResponse{}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(accessResp)
	if err != nil {
		return access, err
	}

	return accessResp.toAccess(), nil
}
