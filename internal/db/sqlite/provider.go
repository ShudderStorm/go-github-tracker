package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Provider struct {
	connPath string
}

func (p *Provider) Open() gorm.Dialector {
	return sqlite.Open(p.connPath)
}

func Memory() *Provider {
	return &Provider{
		connPath: ":memory:",
	}
}

func File(path string) *Provider {
	return &Provider{
		connPath: path,
	}
}
