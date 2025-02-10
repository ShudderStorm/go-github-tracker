package github

import "time"

type User struct {
	ID    uint64 `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Name  string `json:"name"`

	ProfileURL string `json:"html_url"`
	AvatarURL  string `json:"avatar_url"`
}

type Repo struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`

	ProfileURL  string `json:"html_url"`
	Description string `json:"description"`

	Size     uint32 `json:"size"`
	Language string `json:"language"`

	CreationTime   time.Time `json:"created_at"`
	LastUpdateTime time.Time `json:"updated_at"`
	LastPushTime   time.Time `json:"pushed_at"`

	Watchers   int `json:"watchers_count"`
	Stargazers int `json:"stargazers_count"`
}
