package db

import "gorm.io/gorm"

type Model interface {
	Access | User | Repo
}

type crud[M Model] struct {
	model M
	db    *gorm.DB
}

type (
	AccessCrud crud[Access]
	UserCrud   crud[User]
	RepoCrud   crud[Repo]
)

func (crud *crud[M]) Create(value M) error {
	return crud.db.Create(&value).Error
}

func (crud *crud[M]) Read(id uint64) (M, error) {
	model := crud.model
	result := crud.db.First(&model, id)
	return model, result.Error
}

func (crud *crud[M]) Update(id uint64, update M) error {
	return crud.db.Model(&crud.model).Where(id).Updates(update).Error
}

func (crud *crud[M]) Delete(id uint64) error {
	return crud.db.Delete(&crud.model, id).Error
}
