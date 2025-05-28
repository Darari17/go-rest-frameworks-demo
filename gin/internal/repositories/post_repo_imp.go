package repositories

import (
	"errors"
	"fmt"

	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/models"
)

// CreatePost implements PostRepo.
func (p *postRepo) CreatePost(post *models.Post) (*models.Post, error) {
	if err := p.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePostByPostID implements PostRepo.
func (p *postRepo) DeletePostByPostID(postID uint) error {
	res := p.db.Delete(&models.Post{}, postID)
	if res.Error != nil {
		return fmt.Errorf("database error: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}

// GetPostByPostID implements PostRepo.
func (p *postRepo) GetPostByPostID(postID uint) (*models.Post, error) {
	panic("unimplemented")
}

// GetPostsByUserID implements PostRepo.
func (p *postRepo) GetPostsByUserID(userID uint) ([]*models.Post, error) {
	panic("unimplemented")
}

// UpdatePost implements PostRepo.
func (p *postRepo) UpdatePost(post *models.Post) (*models.Post, error) {
	panic("unimplemented")
}
