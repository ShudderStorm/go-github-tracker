package model

type Access struct {
	ID uint64 `gorm:"primaryKey;autoIncrement:false"`

	Token string
	Type  string

	Scopes []Scope `gorm:"foreignKey:AccessID;constraint:OnDelete:CASCADE;"`
	User   User    `gorm:"foreignKey:AccessID;constraint:OnDelete:CASCADE;"`
}

type Scope struct {
	AccessID uint64 `gorm:"primaryKey;autoIncrement:false"`
	Scope    string `gorm:"primaryKey"`
}

type User struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:false"`
	AccessID uint64 `gorm:"unique"`

	Login string
	Name  string
	Email string
}
