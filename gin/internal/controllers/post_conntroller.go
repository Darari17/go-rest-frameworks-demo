package controllers

import (
	"net/http"
	"strconv"

	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/helpers"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/services"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(ps services.PostService) *PostController {
	return &PostController{
		postService: ps,
	}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {

	userId, err := helpers.GetUserFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
		return
	}

	var payload dtos.CreatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
		return
	}

	payload.UserID = userId

	post, err := pc.postService.CreatePost(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Creating post failed: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   post,
	})
}

func (pc *PostController) DeletePost(ctx *gin.Context) {

	userId, err := helpers.GetUserFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
		return
	}

	post, err := pc.postService.GetPostByPostID(uint(postId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to get post: " + err.Error(),
		})
		return
	}

	if post.UserID != userId {
		ctx.AbortWithStatusJSON(http.StatusForbidden, dtos.Response[string]{
			Code:   http.StatusForbidden,
			Status: http.StatusText(http.StatusForbidden),
			Error:  "Forbidden: you are not allowed to delete this post",
		})
		return
	}

	if err := pc.postService.DeletePost(post.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to delete post: " + err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {

	userId, err := helpers.GetUserFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid post ID: " + err.Error(),
		})
		return
	}

	var payload dtos.UpdatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
		return
	}

	payload.ID = uint(postId)
	payload.UserID = userId

	post, err := pc.postService.UpdatePost(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to update post: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   post,
	})
}

func (pc *PostController) GetPostByPostID(ctx *gin.Context)  {}
func (pc *PostController) GetPostsByUserID(ctx *gin.Context) {}
