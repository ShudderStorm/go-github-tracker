package db

import "time"

type Access struct {
	ID uint64 `gorm:"primaryKey"`

	Token  string
	Type   string
	Scopes string
}

type User struct {
	ID uint64 `gorm:"primaryKey"`

	Login string
	Name  string
	Email string

	GitHubURL string
	AvatarURL string

	Access Access `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE"`
	Repos  []Repo `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
}

type Repo struct {
	ID uint64 `gorm:"primaryKey"`

	OwnerID  uint64
	Name     string
	FullName string
	Private  bool

	GitHubURL   string
	Description string
	Watchers    uint
	Stargazers  uint
	Size        uint
	Language    string

	CreatedAt     time.Time
	LastUpdatedAt time.Time
	LastPushedAt  time.Time
}
