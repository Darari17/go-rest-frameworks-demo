package usecases

import (
	"errors"
	"time"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/helper"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/models"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/repositories"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/util"
)

type UserUsecase struct {
	userRepo   repositories.UserRepo
	jwtHandler util.JwtHandler
}

func NewUserUsecase(userRepo repositories.UserRepo, jwtHandler util.JwtHandler) *UserUsecase {
	return &UserUsecase{
		userRepo:   userRepo,
		jwtHandler: jwtHandler,
	}
}

func (u *UserUsecase) Register(req dtos.RegisterRequest) (dtos.UserResponse, error) {

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	user, err := u.userRepo.Create(&models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return dtos.UserResponse{}, err
	}

	token, err := u.jwtHandler.GenerateToken(user.ID)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	response := dtos.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

func (u *UserUsecase) Login(req dtos.LoginRequest) (dtos.UserResponse, error) {

	user, err := u.userRepo.GetEmailOrUsername(req.EmailOrUsername)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	if !helper.CheckPassword(user.Password, req.Password) {
		return dtos.UserResponse{}, errors.New("invalid credentials")
	}

	token, err := u.jwtHandler.GenerateToken(user.ID)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	response := dtos.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}
