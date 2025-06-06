package repositories

import (
	"errors"
	"time"

	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (pr *PostRepository) CreatePost(post *models.Post) (*models.Post, error) {

	if err := pr.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (pr *PostRepository) DeletePostByPostID(postID uint) error {

	result := pr.db.Delete(&models.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}

func (pr *PostRepository) GetPostByPostID(postID uint) (*models.Post, error) {

	var post models.Post
	if err := pr.db.Preload("User").First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (pr *PostRepository) GetPostsByUserID(userID uint) ([]*models.Post, error) {

	var posts []*models.Post
	if err := pr.db.Preload("User").Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *PostRepository) UpdatePost(post *models.Post) (*models.Post, error) {

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
