package services

import (
	"errors"

	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/helpers"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/jwt"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/models"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/repositories"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepo   repositories.UserRepo
	jwtHandler jwt.JWTHandler
}

type UserService interface {
	Login(req dtos.LoginRequest) (string, error)
	Register(req dtos.RegisterRequest) (string, error)
}

func NewUserService(userRepo repositories.UserRepo, jwtHandler jwt.JWTHandler) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtHandler: jwtHandler,
	}
}

// Login implements UserService.
func (u *userService) Login(req dtos.LoginRequest) (string, error) {
	if err := validator.New().Struct(&req); err != nil {
		return "", err
	}

	user, err := u.userRepo.GetUserByEmailOrUsername(req.EmailOrUsername)
	if err != nil {
		return "", err
	}

	if !helpers.CheckPassword(user.Password, req.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := u.jwtHandler.CreateTokenJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Register implements UserService.
func (u *userService) Register(req dtos.RegisterRequest) (string, error) {
	if err := validator.New().Struct(&req); err != nil {
		return "", err
	}

	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	user, err := u.userRepo.CreateUser(&models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	})

	if err != nil {
		return "", err
	}

	token, err := u.jwtHandler.CreateTokenJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
