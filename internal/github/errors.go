package github

import (
	"fmt"
	"net/http"
)

type Error string

const (
	RequestError Error = "request error"
	ParsingError Error = "response parsing error"

	Unathorized Error = "API requires authorization"
	Forbidden   Error = "access forbidden"
)

func (err Error) Error() string {
	return fmt.Sprintf("GitHub API error: %v", string(err))
}

func handleStatusCode(statusCode int) error {
	switch statusCode {
	case http.StatusUnauthorized:
		return fmt.Errorf("%w (http status code: %v)", Unathorized, http.StatusUnauthorized)
	case http.StatusForbidden:
		return fmt.Errorf("%w (http status code: %v)", Forbidden, http.StatusForbidden)
	}
	return nil
}
