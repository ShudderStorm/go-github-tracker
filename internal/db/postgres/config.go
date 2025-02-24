package postgres

import "fmt"

type SSLMode string

const (
	SSLDisable SSLMode = "disable"
	SSLRequire SSLMode = "require"
)

type IConfig interface {
	SetHost(string) IHostConfig
	SetUser(string) IUserConfig

	SetSSLMode(SSLMode) IConfig
	SetDBName(string) IConfig

	Compile() *Provider
}

type IUserConfig interface {
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

	sslmode SSLMode
}

func Config() IConfig {
	return &config{}
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

func (c *config) SetSSLMode(sslmode SSLMode) IConfig {
	c.sslmode = sslmode
	return c
}

func (c *config) SetDBName(dbname string) IConfig {
	c.dbname = dbname
	return c
}

func (c *config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=%s dbname=%s",
		c.host, c.port, c.user, c.password, c.sslmode, c.dbname,
	)
}

func (c *config) Compile() *Provider {
	return &Provider{dsn: c.GetDSN()}
}
