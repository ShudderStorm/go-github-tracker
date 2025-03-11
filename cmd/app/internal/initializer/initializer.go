package initializer

import (
	"github.com/ShudderStorm/go-github-tracker/cmd/app/internal/config"
	"github.com/ShudderStorm/go-github-tracker/internal/cache"
	"github.com/ShudderStorm/go-github-tracker/internal/cache/redis"
	"github.com/ShudderStorm/go-github-tracker/internal/db"
	"github.com/ShudderStorm/go-github-tracker/internal/db/postgres"
	"github.com/ShudderStorm/go-github-tracker/internal/github"
	"github.com/ShudderStorm/go-github-tracker/pkg/oauth"
)

type DatabaseConfig interface {
	ConfigDatabase() config.Database
}

type CacheConfig interface {
	ConfigCache() config.Cache
}

type OAuthConfig interface {
	ConfigOAuth() config.OAuth
}

type Initializer struct {
	dbConfig    DatabaseConfig
	cacheConfig CacheConfig
	oauthConfig OAuthConfig
}

type Option func(*Initializer)

func New(options ...Option) *Initializer {
	init := &Initializer{}
	for _, option := range options {
		option(init)
	}
	return init
}

func WithDatabase(dbConfig DatabaseConfig) Option {
	return func(i *Initializer) {
		i.dbConfig = dbConfig
	}
}

func WithCache(cacheConfig CacheConfig) Option {
	return func(i *Initializer) {
		i.cacheConfig = cacheConfig
	}
}

func WithOAuth(oauthConfig OAuthConfig) Option {
	return func(i *Initializer) {
		i.oauthConfig = oauthConfig
	}
}

func (i Initializer) InitAccessCache() *cache.Cache[cache.Access] {
	cfg := i.cacheConfig.ConfigCache()
	return cache.New(
		redis.New(cfg.Host, cfg.Port),
		cache.SerializeAccess,
		cache.DeserializeAccess,
	)
}

func (i Initializer) InitStateCache() *cache.Cache[cache.State] {
	cfg := i.cacheConfig.ConfigCache()
	return cache.New(
		redis.New(cfg.Host, cfg.Port),
		cache.SerializeState,
		cache.DeserializeState,
	)
}

func (i Initializer) InitOAuth(opts ...oauth.Option) *oauth.Client {
	cfg := i.oauthConfig.ConfigOAuth()
	return oauth.New(
		github.AuthURL,
		github.AccessURL,
		cfg.ClientID,
		cfg.Secret,
		opts...,
	)
}

func (i Initializer) InitDatabase() (*db.Storage, error) {
	cfg := i.dbConfig.ConfigDatabase()
	return db.New(postgres.New(
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	).Compile())
}
