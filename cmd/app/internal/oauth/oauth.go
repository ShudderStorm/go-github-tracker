package oauth

import (
	"github.com/ShudderStorm/go-github-tracker/cmd/app/internal/redis"
	"github.com/ShudderStorm/go-github-tracker/internal/github"
	"github.com/ShudderStorm/go-github-tracker/pkg/oauth"
	"os"
)

const (
	IdEnvKey     string = "OAUTH_ID"
	SecretEnvKey string = "OAUTH_SECRET"
)

var Client *oauth.Client

func init() {
	id := os.Getenv(IdEnvKey)
	secret := os.Getenv(SecretEnvKey)

	Client = oauth.New(
		github.AuthURL,
		github.AccessURL,
		id, secret,
		oauth.WithState(
			NewRedisStateHolder(redis.Storage),
		),
		oauth.WithStateTTL(
			oauth.DefaultStateTTL,
		),
		oauth.WithScopes(
			github.UserScope,
			github.RepoScope,
		),
	)
}

func SetRedirectUrl(url string) {
	oauth.WithRedirectUri(url)(Client)
}
