package db

import (
	"github.com/ShudderStorm/go-github-tracker/internal/db/model"
	"gorm.io/gorm"
)

type AccessCrud struct {
	db *gorm.DB
}

func (crud *AccessCrud) Create(access model.Access) error {
	return crud.db.
		Preload("Scopes").
		Preload("User").
		Create(&access).
		Error
}

func (crud *AccessCrud) Read(id uint64) (model.Access, error) {
	access := model.Access{ID: id}

	err := crud.db.Model(&access).
		Preload("Scopes").
		Preload("User").
		First(&access).
		Error

	return access, err
}

func (crud *AccessCrud) Delete(id uint64) error {
	return crud.db.Delete(&model.Access{ID: id}).Error
}
