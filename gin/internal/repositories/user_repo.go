package repositories

import (
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// Create implements UserRepo.
func (u *UserRepo) Create(user *models.User) (*models.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetEmailOrUsername implements UserRepo.
func (u *UserRepo) GetEmailOrUsername(req string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ? OR username = ?", req, req).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
