package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/helper"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/usecases"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
	validator   *validator.Validate
}

func NewUserHandler(uc usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uc,
		validator:   validator.New(),
	}
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var payload dtos.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[any]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := uh.validator.Struct(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[any]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Validation failed: " + err.Error(),
		})
		return
	}

	user, err := uh.userUsecase.Register(payload)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[any]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Registration failed: " + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dtos.Response[dtos.UserResponse]{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   user,
	})
}

func (uh *UserHandler) Logi(w http.ResponseWriter, r *http.Request) {

	var payload dtos.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[any]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := uh.validator.Struct(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[any]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Validation failed: " + err.Error(),
		})
		return
	}

	user, err := uh.userUsecase.Login(payload)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[any]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Login failed: " + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtos.Response[dtos.UserResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	})
}
