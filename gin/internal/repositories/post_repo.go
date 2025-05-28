package repositories

import (
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/models"
	"gorm.io/gorm"
)

type PostRepo interface {
	CreatePost(post *models.Post) (*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	DeletePostByPostID(postID uint) error

	GetPostByPostID(postID uint) (*models.Post, error)
	GetPostsByUserID(userID uint) ([]*models.Post, error)
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepo{
		db: db,
	}
}
