package repositories

import (
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *models.User) (*models.User, error)
	GetEmailOrUsername(req string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
