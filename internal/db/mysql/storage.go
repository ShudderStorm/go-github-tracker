package mysql

import "gorm.io/gorm"

type Storage struct {
	db *gorm.DB
}
