package oauth

import "fmt"

type Error string

const (
	InternalError Error = "internal error"
	ConfigError   Error = "client configuration error"
	RequestError  Error = "request error"
	ClientError   Error = "http client error"
	ServerError   Error = "http server error"
	ParsingError  Error = "response parsing error"
)

func (err Error) Error() string {
	return fmt.Sprintf("oauth 2.0: %v", string(err))
}
