package services

import (
	"errors"

	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/jwt"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/models"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/repositories"
	"github.com/go-playground/validator/v10"
)

type PostService struct {
	postRepo   repositories.PostRepository
	jwtHandler jwt.JWTHandler
}

func NewPostService(postRepo repositories.PostRepository, jwtHandler jwt.JWTHandler) *PostService {
	return &PostService{
		postRepo:   postRepo,
		jwtHandler: jwtHandler,
	}
}

func (ps *PostService) CreatePost(req dtos.CreatePost) (dtos.PostResponse, error) {

	if err := validator.New().Struct(&req); err != nil {
		return dtos.PostResponse{}, err
	}

	newPost, err := ps.postRepo.CreatePost(&models.Post{
		UserID:   req.UserID,
		Content:  req.Content,
		ImageURL: req.ImageURL,
	})

	if err != nil {
		return dtos.PostResponse{}, err
	}

	response := dtos.PostResponse{
		ID:        newPost.ID,
		UserID:    newPost.UserID,
		Content:   newPost.Content,
		ImageURL:  newPost.ImageURL,
		CreatedAt: newPost.CreatedAt,
		UpdatedAt: newPost.UpdatedAt,
	}

	return response, nil
}

func (ps *PostService) DeletePost(postID uint) error {
	return ps.postRepo.DeletePostByPostID(postID)
}

func (ps *PostService) UpdatePost(req dtos.UpdatePost) (dtos.PostResponse, error) {

	if err := validator.New().Struct(&req); err != nil {
		return dtos.PostResponse{}, err
	}

	existingPost, err := ps.postRepo.GetPostByPostID(req.ID)
	if err != nil {
		return dtos.PostResponse{}, err
	}

	if req.UserID != existingPost.UserID {
		return dtos.PostResponse{}, errors.New("unauthorized to update this post")
	}

	if req.Content != nil {
		existingPost.Content = *req.Content
	}

	if req.ImageURL != nil {
		existingPost.ImageURL = req.ImageURL
	}

	updatedPost, err := ps.postRepo.UpdatePost(existingPost)
	if err != nil {
		return dtos.PostResponse{}, err
	}

	response := dtos.PostResponse{
		ID:        updatedPost.ID,
		UserID:    updatedPost.UserID,
		Content:   updatedPost.Content,
		ImageURL:  updatedPost.ImageURL,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: updatedPost.UpdatedAt,
	}

	return response, nil
}

func (ps *PostService) GetPostByPostID(postID uint) (dtos.PostResponse, error) {
	post, err := ps.postRepo.GetPostByPostID(postID)
	if err != nil {
		return dtos.PostResponse{}, err
	}

	response := dtos.PostResponse{
		ID:        post.ID,
		UserID:    post.UserID,
		Content:   post.Content,
		ImageURL:  post.ImageURL,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return response, nil
}

func (ps *PostService) GetPostsByUserID(userID uint) ([]dtos.PostResponse, error) {
	posts, err := ps.postRepo.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []dtos.PostResponse

	for _, post := range posts {
		responses = append(responses, dtos.PostResponse{
			ID:        post.ID,
			UserID:    post.UserID,
			Content:   post.Content,
			ImageURL:  post.ImageURL,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return responses, nil
}
