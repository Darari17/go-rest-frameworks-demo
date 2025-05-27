package controllers

import (
	"net/http"
	"strconv"

	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/helpers"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/services"
	"github.com/gofiber/fiber/v2"
)

type postController struct {
	postService services.PostService
}

type PostController interface {
	CreatePostController(ctx *fiber.Ctx) error
	GetPostByIdController(ctx *fiber.Ctx) error
	GetPostsController(ctx *fiber.Ctx) error
	UpdatePostController(ctx *fiber.Ctx) error
	DeletePostController(ctx *fiber.Ctx) error
}

func NewPostController(postService services.PostService) PostController {
	return &postController{
		postService: postService,
	}
}

// CreatePostController implements PostController.
func (p *postController) CreatePostController(ctx *fiber.Ctx) error {

	userID, err := helpers.GetUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusUnauthorized,
			Status: http.StatusText(fiber.StatusUnauthorized),
			Error:  "Unauthorized access: " + err.Error(),
		})
	}

	var request dtos.CreatePost
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	if request.UserID != userID {
		return ctx.Status(fiber.StatusForbidden).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusForbidden,
			Status: http.StatusText(fiber.StatusForbidden),
			Error:  "You are not allowed to create post for another user",
		})
	}

	post, err := p.postService.CreatePostService(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Error:  "Failed to create post: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dtos.SuccessResponse{
		Code:   fiber.StatusCreated,
		Status: http.StatusText(fiber.StatusCreated),
		Data:   post,
	})
}

// DeletePostController implements PostController.
func (p *postController) DeletePostController(ctx *fiber.Ctx) error {

	userID, err := helpers.GetUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusUnauthorized,
			Status: http.StatusText(fiber.StatusUnauthorized),
			Error:  "Unauthorized access: " + err.Error(),
		})
	}

	postID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
	}

	post, err := p.postService.FindByPostIDService(uint(postID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Error:  "Failed to get post: " + err.Error(),
		})
	}

	if post.UserID != userID {
		return ctx.Status(fiber.StatusForbidden).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusForbidden,
			Status: http.StatusText(fiber.StatusForbidden),
			Error:  "Forbidden: you are not allowed to delete this post",
		})
	}

	if err := p.postService.DeletePostService(post.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Error:  "Failed to delete post: " + err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetPostByIdController implements PostController.
func (p *postController) GetPostByIdController(ctx *fiber.Ctx) error {

	postID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
	}

	post, err := p.postService.FindByPostIDService(uint(postID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Error:  "Failed to get post: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dtos.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Data:   post,
	})
}

// GetPostsController implements PostController. // Public Posts
func (p *postController) GetPostsController(ctx *fiber.Ctx) error {
	posts, err := p.postService.FindAllPostService()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Error:  "Failed to get posts: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dtos.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Data:   posts,
	})
}

// UpdatePostController implements PostController.
func (p *postController) UpdatePostController(ctx *fiber.Ctx) error {

	userID, err := helpers.GetUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusUnauthorized,
			Status: http.StatusText(fiber.StatusUnauthorized),
			Error:  "Unauthorized access: " + err.Error(),
		})
	}

	postID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
	}

	var request dtos.UpdatePost
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	request.ID = uint(postID)
	request.UserID = userID

	post, err := p.postService.UpdatePostService(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Error:  "Failed to update post: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dtos.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Data:   post,
	})
}
