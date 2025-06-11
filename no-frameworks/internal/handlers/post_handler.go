package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/helper"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/usecases"
	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	postUsecase usecases.PostUsecase
	validator   *validator.Validate
}

func NewPostHandler(postUsecase usecases.PostUsecase) *PostHandler {
	return &PostHandler{
		postUsecase: postUsecase,
		validator:   validator.New(),
	}
}

func (ph *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		helper.JSON(w, http.StatusMethodNotAllowed, dtos.Response[any]{
			Code:   http.StatusMethodNotAllowed,
			Status: http.StatusText(http.StatusMethodNotAllowed),
			Error:  "Method not allowed",
		})
		return
	}

	ctx := r.Context()

	userId, err := helper.GetUserFromContext(ctx)
	if err != nil {
		helper.JSON(w, http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
		return
	}

	var payload dtos.CreatePost
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := ph.validator.Struct(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Validation failed: " + err.Error(),
		})
		return
	}

	payload.UserID = userId
	post, err := ph.postUsecase.CreatePost(payload)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Creating post failed: " + err.Error(),
		})
		return
	}

	helper.JSON(w, http.StatusCreated, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   post,
	})
}

func (ph *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		helper.JSON(w, http.StatusMethodNotAllowed, dtos.Response[any]{
			Code:   http.StatusMethodNotAllowed,
			Status: http.StatusText(http.StatusMethodNotAllowed),
			Error:  "Method not allowed",
		})
		return
	}

	ctx := r.Context()

	userId, err := helper.GetUserFromContext(ctx)
	if err != nil {
		helper.JSON(w, http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
		return
	}

	postId, err := helper.ExtractPathID(r)
	if err != nil || postId <= 0 {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  err.Error(),
		})
		return
	}

	post, err := ph.postUsecase.FindPostByPostId(uint(postId))
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to get post: " + err.Error(),
		})
		return
	}

	if post.UserID != userId {
		helper.JSON(w, http.StatusForbidden, dtos.Response[string]{
			Code:   http.StatusForbidden,
			Status: http.StatusText(http.StatusForbidden),
			Error:  "Forbidden: you are not allowed to delete this post",
		})
		return
	}

	if err := ph.postUsecase.DeletePost(post.ID); err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to delete post: " + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ph *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		helper.JSON(w, http.StatusMethodNotAllowed, dtos.Response[any]{
			Code:   http.StatusMethodNotAllowed,
			Status: http.StatusText(http.StatusMethodNotAllowed),
			Error:  "Method not allowed",
		})
		return
	}

	ctx := r.Context()

	userId, err := helper.GetUserFromContext(ctx)
	if err != nil {
		helper.JSON(w, http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
		return
	}

	postId, err := helper.ExtractPathID(r)
	if err != nil || postId <= 0 {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  err.Error(),
		})
		return
	}

	var payload dtos.UpdatePost
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := ph.validator.Struct(&payload); err != nil {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Validation failed: " + err.Error(),
		})
		return
	}

	payload.ID = uint(postId)
	payload.UserID = userId

	post, err := ph.postUsecase.UpdatePost(payload)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to update post: " + err.Error(),
		})
		return
	}

	helper.JSON(w, http.StatusOK, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   post,
	})
}

func (ph *PostHandler) FindPostByPostId(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		helper.JSON(w, http.StatusMethodNotAllowed, dtos.Response[any]{
			Code:   http.StatusMethodNotAllowed,
			Status: http.StatusText(http.StatusMethodNotAllowed),
			Error:  "Method not allowed",
		})
		return
	}

	postId, err := helper.ExtractPathID(r)
	if err != nil || postId <= 0 {
		helper.JSON(w, http.StatusBadRequest, dtos.Response[string]{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  err.Error(),
		})
		return
	}

	post, err := ph.postUsecase.FindPostByPostId(uint(postId))
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to get post: " + err.Error(),
		})
		return
	}

	helper.JSON(w, http.StatusOK, dtos.Response[dtos.PostResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   post,
	})
}

func (ph *PostHandler) FindPostsByUserId(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		helper.JSON(w, http.StatusMethodNotAllowed, dtos.Response[any]{
			Code:   http.StatusMethodNotAllowed,
			Status: http.StatusText(http.StatusMethodNotAllowed),
			Error:  "Method not allowed",
		})
		return
	}

	ctx := r.Context()

	userId, err := helper.GetUserFromContext(ctx)
	if err != nil {
		helper.JSON(w, http.StatusUnauthorized, dtos.Response[string]{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  "Unauthorized: " + err.Error(),
		})
		return
	}

	posts, err := ph.postUsecase.FindPostsByUserId(userId)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, dtos.Response[string]{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Error:  "Failed to get posts: " + err.Error(),
		})
		return
	}

	helper.JSON(w, http.StatusOK, dtos.Response[[]dtos.PostResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   posts,
	})
}
