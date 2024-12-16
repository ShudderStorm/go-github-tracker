package oauth

import "fmt"

type Error string

const (
	AccessDenied Error = "authorization request was denied"
	ExpiredToken Error = "device code has expired"
)

func (err Error) Error() string {
	return fmt.Sprintf("oauth 2.0: %v", string(err))
}
