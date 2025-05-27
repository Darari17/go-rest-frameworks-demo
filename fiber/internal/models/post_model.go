package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	ImageURL  *string   `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
