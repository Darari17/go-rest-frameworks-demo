package controllers

import (
	"net/http"
	"strconv"

	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/helpers"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/services"
	"github.com/labstack/echo/v4"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(ps services.PostService) *PostController {
	return &PostController{
		postService: ps,
	}
}

func (pc *PostController) CreatePost(c echo.Context) error {

	userId, err := helpers.GetUserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
	}

	var payload dtos.CreatePost
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	payload.UserID = userId

	post, err := pc.postService.CreatePost(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Creating post failed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   post,
	})
}

func (pc *PostController) DeletePost(c echo.Context) error {

	userId, err := helpers.GetUserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
	}

	post, err := pc.postService.GetPostByPostID(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to get post: " + err.Error(),
		})
	}

	if post.UserID != userId {
		return c.JSON(http.StatusForbidden, dtos.Response[string]{
			Code:   http.StatusForbidden,
			Status: http.StatusText(http.StatusForbidden),
			Error:  "Forbidden: you are not allowed to delete this post",
		})
	}

	if err := pc.postService.DeletePost(post.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to delete post: " + err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (pc *PostController) UpdatePost(c echo.Context) error {

	userId, err := helpers.GetUserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
	}

	var payload dtos.UpdatePost
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
	}

	payload.ID = uint(postId)
	payload.UserID = userId

	post, err := pc.postService.UpdatePost(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to update post: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   post,
	})
}

func (pc *PostController) GetPostByPostID(c echo.Context) error {

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		return c.JSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
	}

	post, err := pc.postService.GetPostByPostID(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to get post: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   post,
	})
}

func (pc *PostController) GetPostsByUserID(c echo.Context) error {

	userId, err := helpers.GetUserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
	}

	posts, err := pc.postService.GetPostsByUserID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to get posts: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.Response[[]dtos.PostResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   posts,
	})
}
