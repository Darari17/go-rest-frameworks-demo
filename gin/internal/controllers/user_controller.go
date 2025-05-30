package controllers

import (
	"net/http"

	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {

	var payload dtos.RegisterRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "invalid request body: " + err.Error(),
		})
		return
	}

	user, err := uc.userService.Register(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Registration failed: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.Response[dtos.UserResponse]{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   user,
	})
}

func (uc *UserController) Login(ctx *gin.Context) {

	var payload dtos.LoginRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "invalid request body: " + err.Error(),
		})
		return
	}

	user, err := uc.userService.Login(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Login failed: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.Response[dtos.UserResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	})
}
