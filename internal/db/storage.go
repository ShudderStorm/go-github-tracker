package db

import (
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

type Provider interface {
	Open() gorm.Dialector
}

func New(dbprovider Provider) (*Storage, error) {
	db, err := gorm.Open(dbprovider.Open(), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Migrate() error {
	return s.db.AutoMigrate(&Access{}, &User{}, &UserProfile{})
}
