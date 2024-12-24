package github

import "github.com/ShudderStorm/go-github-tracker/pkg/oauth"

const (
	AuthURL   string = "https://github.com/login/oauth/authorize"
	AccessURL string = "https://github.com/login/oauth/access_token"
)

const (
	RepoScope string = "repo"
	UserScope string = "user"
)

func NewOAuthClient(id string, secret string, opts ...oauth.Option) *oauth.Client {
	return oauth.New(AuthURL, AccessURL, id, secret, opts...)
}
