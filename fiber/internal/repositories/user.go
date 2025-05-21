package repositories

import (
	"github.com/go-rest-frameworks-demo/fiber/internal/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type UserRepo interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUserName(username string) (models.User, error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

// CreateUser implements UserRepo.
func (u *userRepo) CreateUser(user models.User) (models.User, error) {
	if err := u.db.Model(&models.User{}).Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserByEmail implements UserRepo.
func (u *userRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := u.db.Model(&models.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserByUserName implements UserRepo.
func (u *userRepo) GetUserByUserName(username string) (models.User, error) {
	var user models.User
	if err := u.db.Model(&models.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
