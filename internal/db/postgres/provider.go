package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dsnProvider interface {
	dsn() string
}

type Provider struct {
	dsn string
}

func (p *Provider) Open() gorm.Dialector {
	return postgres.Open(p.dsn)
}
