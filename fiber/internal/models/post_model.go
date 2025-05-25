package models

import "time"

type Post struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	UserID    uint    `gorm:"not null"`
	Content   string  `gorm:"type:text;not null"`
	ImageURL  *string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User `gorm:"foreignKey:UserID;references:ID"`
}
