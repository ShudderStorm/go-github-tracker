package github

import (
	"github.com/ShudderStorm/go-github-tracker/pkg/oauth"
	"github.com/go-resty/resty/v2"
)

type RestApi struct {
	access oauth.Access
	client resty.Client
}

func NewRestApiClient(access oauth.Access) *RestApi {
	return &RestApi{
		access: access,
		client: *resty.New(),
	}
}
