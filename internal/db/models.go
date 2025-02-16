package db

type Access struct {
	ID uint `gorm:"primaryKey"`

	Token  string
	Type   string
	Scopes string

	User User `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	ID uint64 `gorm:"primaryKey"`

	Login string

	Profile Profile `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE"`
}

type Profile struct {
	ID uint `gorm:"primaryKey"`

	Name  string
	Email string

	GitHubURL string
	AvatarURL string
}
