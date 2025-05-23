package repositories

import (
	"github.com/go-rest-frameworks-demo/fiber/internal/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmailOrUsername(req string) (*models.User, error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

// CreateUser implements UserRepo.
func (u *userRepo) CreateUser(user *models.User) (*models.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail implements UserRepo.
func (u *userRepo) GetUserByEmailOrUsername(req string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ? OR username = ?", req, req).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
