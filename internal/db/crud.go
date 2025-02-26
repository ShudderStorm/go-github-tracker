package db

import (
	"github.com/ShudderStorm/go-github-tracker/internal/db/model"
	"gorm.io/gorm"
)

type AccessCrud struct {
	db *gorm.DB
}

func (crud *AccessCrud) Create(access model.Access) error {
	err := crud.db.Create(&access).Error
	if err != nil {
		return err
	}

	for _, scope := range access.Scopes {
		err = crud.db.Create(&scope).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (crud *AccessCrud) Read(id uint64) (model.Access, error) {
	access := model.Access{ID: id}
	err := crud.db.First(&access).Error
	return access, err
}

type UserCrud struct {
	db *gorm.DB
}

func (crud *UserCrud) Create(user model.User) error {
	return crud.db.Create(&user).Error
}

func (crud *UserCrud) Read(id uint64) (model.User, error) {
	user := model.User{ID: id}
	err := crud.db.First(&user).Error
	return user, err
}

func (crud *UserCrud) GetOwnedRepos(id uint64) ([]model.Repo, error) {
	var repos []model.Repo
	err := crud.db.Where(&model.Repo{OwnerID: id}).Find(&repos).Error
	return repos, err
}
