package oauth

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParseJSONTestHolder[S oauthStatus] struct {
	name        string
	reader      io.Reader
	expected    S
	expectedErr error
}

func TestParseJSON(t *testing.T) {
	t.Parallel()
	verificationTests := []ParseJSONTestHolder[Verification]{
		{
			"Verification::Valid_JSON",
			bytes.NewReader([]byte(`{
				"device_code": "3584d83530557fdd1f46af8289938c8ef79f9dc5",
				"user_code": "WDJB-MJHT",
				"verification_uri": "https://github.com/login/device",
				"expires_in": 900,
				"interval": 5
			}`)),
			Verification{
				"3584d83530557fdd1f46af8289938c8ef79f9dc5",
				"WDJB-MJHT",
				"https://github.com/login/device",
				900,
				5,
			},
			nil,
		},
		{
			"Verification::Empty_Reader",
			bytes.NewReader(make([]byte, 0)),
			Verification{},
			io.EOF,
		},
	}

	for _, test := range verificationTests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := ParseJSON[Verification](test.reader)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}

	pollStatusTests := []ParseJSONTestHolder[PollStatus]{
		{
			"PollStatus::Valid_JSON",
			bytes.NewReader([]byte(`{
				"access_token": "gho_16C7e42F292c6912E7710c838347Ae178B4a",
				"token_type": "bearer",
				"scope": "repo,gist"
			}`)),
			PollStatus{
				"gho_16C7e42F292c6912E7710c838347Ae178B4a",
				"bearer",
				"repo,gist",
			},
			nil,
		},
		{
			"PollStatus::Empty_Reader",
			bytes.NewReader(make([]byte, 0)),
			PollStatus{},
			io.EOF,
		},
	}

	for _, test := range pollStatusTests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := ParseJSON[PollStatus](test.reader)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
