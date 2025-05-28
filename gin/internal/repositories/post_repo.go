package repositories

import (
	"errors"

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

// CreatePost implements PostRepo.
func (p *postRepo) CreatePost(post *models.Post) (*models.Post, error) {

	if err := p.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePostByPostID implements PostRepo.
func (p *postRepo) DeletePostByPostID(postID uint) error {

	result := p.db.Delete(&models.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}

// GetPostByPostID implements PostRepo.
func (p *postRepo) GetPostByPostID(postID uint) (*models.Post, error) {

	var post models.Post
	if err := p.db.Preload("User").First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

// GetPostsByUserID implements PostRepo.
func (p *postRepo) GetPostsByUserID(userID uint) ([]*models.Post, error) {

	var posts []*models.Post
	if err := p.db.Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost implements PostRepo.
func (p *postRepo) UpdatePost(post *models.Post) (*models.Post, error) {

	// cek user id dan post id
	var existingPost models.Post
	if err := p.db.Where("id = ? AND user_id = ?", post.ID, post.UserID).First(&existingPost).Error; err != nil {
		return nil, errors.New("post not found")
	}

	// update yang pilih saja
	err := p.db.Model(&existingPost).
		Select("Content", "ImageURL", "UpdatedAt").
		Updates(post).
		Error
	if err != nil {
		return nil, err
	}

	// get data terbaru
	if err := p.db.Preload("User").First(&existingPost, post.ID).Error; err != nil {
		return nil, err
	}

	return &existingPost, nil
}
