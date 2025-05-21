package models

import "time"

type Post struct {
	ID        uint
	UserID    uint
	Content   string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}
