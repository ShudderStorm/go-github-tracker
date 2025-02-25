package postgres

import (
	"fmt"
	"net/url"

	"github.com/ShudderStorm/go-github-tracker/internal/db/postgres/ssl"
)

type IConfig interface {
	SetHost(string) IHostConfig
	SetUser(string) IUserConfig

	SetSSLMode(ssl.Mode) IConfig
	SetDBName(string) IConfig

	Compile() *Provider
}

type IUserConfig interface {
	IConfig
	SetPassword(string) IConfig
}

type IHostConfig interface {
	SetPort(uint16) IConfig
}

type config struct {
	host string
	port uint16

	user     string
	password string
	dbname   string

	params *url.Values
}

func Config() IConfig {
	return &config{params: &url.Values{}}
}

func (c *config) SetHost(host string) IHostConfig {
	c.host = host
	return c
}

func (c *config) SetPort(port uint16) IConfig {
	c.port = port
	return c
}

func (c *config) SetUser(user string) IUserConfig {
	c.user = user
	return c
}

func (c *config) SetPassword(password string) IConfig {
	c.password = password
	return c
}

func (c *config) SetSSLMode(sslmode ssl.Mode) IConfig {
	c.params.Set("sslmode", string(sslmode))
	return c
}

func (c *config) SetDBName(dbname string) IConfig {
	c.dbname = dbname
	return c
}

func (c *config) GetDSN() string {
	dsn := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.user, c.password),
		Host:     fmt.Sprintf("%s:%d", c.host, c.port),
		Path:     c.dbname,
		RawQuery: c.params.Encode(),
	}

	return dsn.String()
}

func (c *config) Compile() *Provider {
	return &Provider{dsn: c.GetDSN()}
}
