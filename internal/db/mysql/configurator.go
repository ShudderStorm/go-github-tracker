package mysql

import "github.com/go-sql-driver/mysql"

type Configurator struct {
	config *mysql.Config
}

func New() *Configurator {
	return &Configurator{
		config: mysql.NewConfig(),
	}
}

func (c *Configurator) FormatDSN() string {
	return c.config.FormatDSN()
}

func (c *Configurator) SetUser(user string) *Configurator {
	c.config.User = user
	return c
}

func (c *Configurator) SetPassword(password string) *Configurator {
	c.config.Passwd = password
	return c
}
