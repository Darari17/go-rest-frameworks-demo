package dtos

import (
	"time"
)

type CreatePost struct {
	UserID   uint    `json:"user_id" validate:"required"`
	Content  string  `json:"content" validate:"required,min=1"`
	ImageURL *string `json:"image_url" validate:"omitempty,url"`
}

type UpdatePost struct {
	ID       uint    `json:"id" validate:"required"`
	UserID   uint    `json:"user_id" validate:"required"`
	Content  *string `json:"content" validate:"omitempty,min=1"`
	ImageURL *string `json:"image_url" validate:"omitempty,url"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	ImageURL  *string   `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
