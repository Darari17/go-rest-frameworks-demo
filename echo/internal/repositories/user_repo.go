package repositories

import (
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {

	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetEmailOrUsername(req string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ? OR username = ?", req, req).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
