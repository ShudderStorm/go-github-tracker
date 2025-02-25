package postgres

import (
	"fmt"
	"github.com/ShudderStorm/go-github-tracker/internal/db/postgres/ssl"
	"net/url"
)

type Config struct {
	host string
	port uint16

	user     string
	password string
	database string

	params *url.Values
}

type ConfigOption func(*Config)

func New(host string, port uint16, user, password, database string, opts ...ConfigOption) *Config {
	cfg := &Config{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
		params:   &url.Values{},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func WithSSL(mode ssl.Mode) ConfigOption {
	return func(cfg *Config) {
		cfg.params.Set("sslmode", string(mode))
	}
}

func (cfg *Config) GetDSN() string {
	dsn := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.user, cfg.password),
		Host:     fmt.Sprintf("%s:%d", cfg.host, cfg.port),
		Path:     cfg.database,
		RawQuery: cfg.params.Encode(),
	}

	return dsn.String()
}

func (cfg *Config) Compile() *Provider {
	return &Provider{dsn: cfg.GetDSN()}
}
