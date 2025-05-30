package services

import (
	"errors"
	"time"

	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/helpers"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/jwt"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/models"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/repositories"
)

type UserService struct {
	userRepo   repositories.UserRepo
	jwtHandler jwt.JWTHandler
}

func NewUserService(ur repositories.UserRepo, jh jwt.JWTHandler) *UserService {
	return &UserService{
		userRepo:   ur,
		jwtHandler: jh,
	}
}

func (us *UserService) Register(req dtos.RegisterRequest) (dtos.UserResponse, error) {

	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	user, err := us.userRepo.Create(&models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return dtos.UserResponse{}, err
	}

	token, err := us.jwtHandler.GenerateToken(user.ID)
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

func (us *UserService) Login(req dtos.LoginRequest) (dtos.UserResponse, error) {

	user, err := us.userRepo.GetEmailOrUsername(req.EmailOrUsername)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	if !helpers.CheckPassword(user.Password, req.Password) {
		return dtos.UserResponse{}, errors.New("invalid credentials")
	}

	token, err := us.jwtHandler.GenerateToken(user.ID)
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
