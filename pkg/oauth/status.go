package oauth

import (
	"encoding/json"
	"io"
)

type oauthStatus interface {
	Verification | PollStatus
}

type Verification struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUri string `json:"verification_uri"`
	TTL             int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

type PollStatus struct {
	АccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func parseJSON[S oauthStatus](reader io.Reader) (S, error) {
	var status S
	decodingError := json.NewDecoder(reader).Decode(&status)
	return status, decodingError
}
