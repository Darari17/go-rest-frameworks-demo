package controllers

import (
	"github.com/go-rest-frameworks-demo/fiber/internal/dtos"
	"github.com/go-rest-frameworks-demo/fiber/internal/services"
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := u.userService.Login(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dtos.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "OK",
		Data:   token,
	})
}

// Register implements UserController.
func (u *userController) Register(ctx *fiber.Ctx) error {

	var request dtos.RegisterRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := u.userService.Register(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dtos.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   token,
	})
}
