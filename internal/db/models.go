package db

import "time"

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

	Profile UserProfile `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Repos   Repo        `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL"`
}

type UserProfile struct {
	UserID uint64

	Name  string
	Email string

	GitHubURL string
	AvatarURL string
}

type Repo struct {
	ID uint64 `gorm:"primaryKey"`

	OwnerID uint64

	Profile RepoProfile `gorm:"foreignKey:RepoID;constraint:OnDelete:CASCADE"`
	Meta    RepoMeta    `gorm:"foreignKey:RepoID;constraint:OnDelete:CASCADE"`
}

type RepoProfile struct {
	RepoID uint64

	Name     string
	FullName string
	Private  bool

	GitHubURL   string
	Description string
}

type RepoMeta struct {
	RepoID uint64

	CreationTime   time.Time
	LastUpdateTime time.Time
	LastPushTime   time.Time

	Size     uint
	Language string

	Watchers   uint
	Stargazers uint
}
