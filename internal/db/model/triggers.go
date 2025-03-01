package model

import "gorm.io/gorm"

func (a *Access) AfterCreate(db *gorm.DB) error {
	return db.Create(&a.Scopes).Error
}
