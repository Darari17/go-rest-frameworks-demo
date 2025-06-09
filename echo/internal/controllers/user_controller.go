package controllers

import (
	"net/http"

	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService services.UserService
	validator   *validator.Validate
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{
		userService: us,
		validator:   validator.New(),
	}
}

func (uc *UserController) Register(c echo.Context) error {

	var payload dtos.RegisterRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	if err := uc.validator.Struct(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	user, err := uc.userService.Register(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Registration failed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dtos.Response[dtos.UserResponse]{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   user,
	})
}

func (uc *UserController) Login(c echo.Context) error {

	var payload dtos.LoginRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	if err := uc.validator.Struct(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	user, err := uc.userService.Login(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Login failed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.Response[dtos.UserResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	})
}
