package repositories

import (
	"errors"
	"time"

	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/models"
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

// CreatePost implements PostRepo.
func (pr *PostRepo) CreatePost(post *models.Post) (*models.Post, error) {

	if err := pr.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePostByPostID implements PostRepo.
func (pr *PostRepo) DeletePostByPostID(postID uint) error {

	result := pr.db.Delete(&models.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}

// GetPostByPostID implements PostRepo.
func (pr *PostRepo) GetPostByPostID(postID uint) (*models.Post, error) {

	var post models.Post
	if err := pr.db.Preload("User").First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

// GetPostsByUserID implements PostRepo.
func (pr *PostRepo) GetPostsByUserID(userID uint) ([]*models.Post, error) {

	var posts []*models.Post
	if err := pr.db.Preload("User").Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost implements PostRepo.
func (pr *PostRepo) UpdatePost(post *models.Post) (*models.Post, error) {

	result := pr.db.Model(&models.Post{}).
		Where("id = ? AND user_id = ?", post.ID, post.UserID).
		Updates(map[string]interface{}{
			"Content":   post.Content,
			"ImageURL":  post.ImageURL,
			"UpdatedAt": time.Now(),
		})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("post not found")
	}

	post.UpdatedAt = time.Now()
	return post, nil
}
