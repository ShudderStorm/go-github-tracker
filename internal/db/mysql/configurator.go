package mysql

import "github.com/go-sql-driver/mysql"

type IConfigurator interface {
	FormatDSN() string
	SetUser(string) IUserConfigurator
	SetAddres(string) IConfigurator
	SetDBName(string) IConfigurator
	SetParams(map[string]string) IConfigurator
}

type IUserConfigurator interface {
	IConfigurator
	SetPassword(string) IConfigurator
}

type configurator struct {
	config *mysql.Config
}

func New() *configurator {
	return &configurator{
		config: mysql.NewConfig(),
	}
}

func (c *configurator) FormatDSN() string {
	return c.config.FormatDSN()
}

func (c *configurator) SetUser(user string) IUserConfigurator {
	c.config.User = user
	return c
}

func (c *configurator) SetPassword(password string) IConfigurator {
	c.config.Passwd = password
	return c
}

func (c *configurator) SetAddres(addres string) IConfigurator {
	c.config.Addr = addres
	return c
}

func (c *configurator) SetDBName(dbname string) IConfigurator {
	c.config.DBName = dbname
	return c
}

func (c *configurator) SetParams(params map[string]string) IConfigurator {
	for param, value := range params {
		c.config.Params[param] = value
	}
	return c
}
