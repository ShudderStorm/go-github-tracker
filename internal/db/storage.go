package db

import "gorm.io/gorm"

type Storage struct {
	db *gorm.DB
}

func New(dbprovider func() (*gorm.DB, error)) (*Storage, error) {
	db, err := dbprovider()

	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Migrate() error {
	return s.db.AutoMigrate(&Access{}, &User{}, &UserProfile{})
}
