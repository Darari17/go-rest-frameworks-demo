package repositories

import (
	"github.com/go-rest-frameworks-demo/fiber/internal/models"
	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

type PostRepo interface {
	CreatePost(post models.Post) (models.Post, error)
	DeletePost(postId uint) error
	GetPostById(postId uint) (models.Post, error)
	GetPosts() ([]models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepo{
		db: db,
	}
}

// CreatePost implements PostRepo.
func (p *postRepo) CreatePost(post models.Post) (models.Post, error) {
	panic("unimplemented")
}

// DeletePost implements PostRepo.
func (p *postRepo) DeletePost(postId uint) error {
	panic("unimplemented")
}

// GetPostById implements PostRepo.
func (p *postRepo) GetPostById(postId uint) (models.Post, error) {
	panic("unimplemented")
}

// GetPosts implements PostRepo.
func (p *postRepo) GetPosts() ([]models.Post, error) {
	panic("unimplemented")
}

// UpdatePost implements PostRepo.
func (p *postRepo) UpdatePost(post models.Post) (models.Post, error) {
	panic("unimplemented")
}
