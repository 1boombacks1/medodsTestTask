package models

import "time"

type Session struct {
	Id           string
	Guid         string
	RefreshToken string
	ExpiresIn    int64
	CreatedAt    time.Time
}
