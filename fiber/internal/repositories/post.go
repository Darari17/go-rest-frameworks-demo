package repositories

import (
	"github.com/go-rest-frameworks-demo/fiber/internal/models"
	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

type PostRepo interface {
	CreatePost(post *models.Post) (*models.Post, error)
	DeletePost(postId uint) error
	GetPostById(postId uint) (*models.Post, error)
	GetPosts() ([]*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepo{
		db: db,
	}
}

// CreatePost implements PostRepo.
func (p *postRepo) CreatePost(post *models.Post) (*models.Post, error) {
	if err := p.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePost implements PostRepo.
func (p *postRepo) DeletePost(postId uint) error {
	return p.db.Delete(&models.Post{}, postId).Error
}

// GetPostById implements PostRepo.
func (p *postRepo) GetPostById(postId uint) (*models.Post, error) {
	var post models.Post
	if err := p.db.First(&post, postId).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPosts implements PostRepo.
func (p *postRepo) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost implements PostRepo.
func (p *postRepo) UpdatePost(post *models.Post) (*models.Post, error) {
	if err := p.db.Model(&models.Post{}).Where("id = ?", post.ID).Updates(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}
