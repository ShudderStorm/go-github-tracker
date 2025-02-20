package mysql

import "github.com/go-sql-driver/mysql"

type IConfig interface {
	FormatDSN() string
	SetUser(string) IUserConfig
	UseTCP() ITCPConfig
	SetDBName(string) IConfig
	AddParams(map[string]string) IConfig
}

type IUserConfig interface {
	IConfig
	SetPassword(string) IConfig
}

type ITCPConfig interface {
	SetAddress(string) IConfig
}

type configurator struct {
	config *mysql.Config
}

func New() IConfig {
	return &configurator{
		config: mysql.NewConfig(),
	}
}

func (c *configurator) FormatDSN() string {
	return c.config.FormatDSN()
}

func (c *configurator) SetUser(user string) IUserConfig {
	c.config.User = user
	return c
}

func (c *configurator) SetPassword(password string) IConfig {
	c.config.Passwd = password
	return c
}

func (c *configurator) UseTCP() ITCPConfig {
	c.config.Net = "tcp"
	return c
}

func (c *configurator) SetAddress(addres string) IConfig {
	c.config.Addr = addres
	return c
}

func (c *configurator) SetDBName(dbname string) IConfig {
	c.config.DBName = dbname
	return c
}

func (c *configurator) AddParams(params map[string]string) IConfig {
	if c.config.Params == nil {
		c.config.Params = params
		return c
	}

	for param, value := range params {
		c.config.Params[param] = value
	}
	return c
}
