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
	params := url.Values{
		"client_id":    {c.config.id},
		"redirect_uri": {c.config.redirectUri},
		"scopes":       c.config.scopes,
	}.Encode()

	return fmt.Sprintf("%s?%s", c.config.authEndpoint, params), nil
}

func (c *Client) Exchange(ctx context.Context, code string) (Access, error) {
	var access Access

	params := url.Values{
		"client_id":     {c.config.id},
		"client_secret": {c.config.secret},
		"code":          {code},
		"redirect_uri":  {c.config.redirectUri},
	}.Encode()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, c.config.tokenEndpoint, strings.NewReader(params),
	)

	if err != nil {
		return access, err
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json"},
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return access, err
	}

	accessResp := &accessResponse{}
	err = json.NewDecoder(resp.Body).Decode(accessResp)
	if err != nil {
		return access, err
	}

	return accessResp.toAccess(), nil
}
