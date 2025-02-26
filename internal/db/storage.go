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

func New(provider Provider) (*Storage, error) {
	db, err := gorm.Open(provider.Open(), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Migrate() error {
	return s.db.AutoMigrate(&Access{}, &User{}, &Repo{})
}

func (s *Storage) Access() *AccessCrud {
	return &AccessCrud{model: Access{}, db: s.db}
}

func (s *Storage) User() *UserCrud {
	return &UserCrud{model: User{}, db: s.db}
}

func (s *Storage) Repo() *RepoCrud {
	return &RepoCrud{model: Repo{}, db: s.db}
}
