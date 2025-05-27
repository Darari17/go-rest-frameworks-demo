package controllers

import (
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/services"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService services.UserService
}

type UserController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// Login implements UserController.
func (u *userController) Login(ctx *fiber.Ctx) error {

	var request dtos.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Error:  "invalid request body: " + err.Error(),
		})
	}

	token, err := u.userService.Login(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Error:  "login failed: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dtos.SuccessResponse{
		Code:   fiber.StatusCreated,
		Status: "OK",
		Data:   token,
	})
}

// Register implements UserController.
func (u *userController) Register(ctx *fiber.Ctx) error {

	var request dtos.RegisterRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Error:  "invalid request body: " + err.Error(),
		})
	}

	token, err := u.userService.Register(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Error:  "registration failed: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dtos.SuccessResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   token,
	})
}
