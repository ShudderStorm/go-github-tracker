package db

import (
	"github.com/ShudderStorm/go-github-tracker/internal/db/model"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB

	Access *AccessCrud
	User   *UserCrud
}

type Provider interface {
	Open() gorm.Dialector
}

func New(provider Provider) (*Storage, error) {
	db, err := gorm.Open(provider.Open(), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Storage{
		db:     db,
		Access: &AccessCrud{db: db},
		User:   &UserCrud{db: db},
	}, nil
}

func (s *Storage) Migrate() error {
	return s.db.AutoMigrate(&model.Access{}, &model.Scope{}, &model.User{}, &model.Repo{})
}
