package repositories

import (
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/models"
	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

type PostRepo interface {
	CreatePostRepo(post *models.Post) (*models.Post, error)
	DeletePostRepo(postId uint) error
	GetPostByIdRepo(postId uint) (*models.Post, error)
	GetPostsRepo() ([]*models.Post, error)
	UpdatePostRepo(post *models.Post) (*models.Post, error)
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepo{
		db: db,
	}
}

// CreatePost implements PostRepo.
func (p *postRepo) CreatePostRepo(post *models.Post) (*models.Post, error) {
	if err := p.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePost implements PostRepo.
func (p *postRepo) DeletePostRepo(postId uint) error {
	return p.db.Delete(&models.Post{}, postId).Error
}

// GetPostById implements PostRepo.
func (p *postRepo) GetPostByIdRepo(postId uint) (*models.Post, error) {
	var post models.Post
	if err := p.db.First(&post, postId).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPosts implements PostRepo.
func (p *postRepo) GetPostsRepo() ([]*models.Post, error) {
	var posts []*models.Post
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost implements PostRepo.
func (p *postRepo) UpdatePostRepo(post *models.Post) (*models.Post, error) {
	if err := p.db.Model(&models.Post{}).Where("id = ?", post.ID).Updates(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}
