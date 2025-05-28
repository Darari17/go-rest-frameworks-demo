package repositories

import "github.com/Darari17/go-rest-frameworks-demo/gin/internal/models"

// Create implements UserRepo.
func (u *userRepo) Create(user *models.User) (*models.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetEmailOrUsername implements UserRepo.
func (u *userRepo) GetEmailOrUsername(req string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ? OR username = ?", req, req).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
