package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/go-rest-frameworks-demo/fiber/internal/dtos"
	"github.com/go-rest-frameworks-demo/fiber/internal/models"
	"github.com/go-rest-frameworks-demo/fiber/internal/repositories"
)

type postService struct {
	postRepo repositories.PostRepo
}

type PostService interface {
	CreatePost(req dtos.CreatePost) (dtos.PostResponse, error)
	FindByPostID(postId uint) (dtos.PostResponse, error)
	FindAllPost() ([]dtos.PostResponse, error)
	UpdatePost(req dtos.UpdatePost) (dtos.PostResponse, error)
	DeletePost(postId uint) error
}

func NewPostService(postRepo repositories.PostRepo) PostService {
	return &postService{
		postRepo: postRepo,
	}
}

// CreatePost implements PostService.
func (p *postService) CreatePost(req dtos.CreatePost) (dtos.PostResponse, error) {
	if err := validator.New().Struct(&req); err != nil {
		return dtos.PostResponse{}, err
	}

	post, err := p.postRepo.CreatePost(&models.Post{
		UserID:   req.UserID,
		Content:  req.Content,
		ImageURL: req.ImageURL,
	})
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

// DeletePost implements PostService.
func (p *postService) DeletePost(postId uint) error {
	return p.postRepo.DeletePost(postId)
}

// FindAllPost implements PostService.
func (p *postService) FindAllPost() ([]dtos.PostResponse, error) {
	posts, err := p.postRepo.GetPosts()
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

// FindByPostID implements PostService.
func (p *postService) FindByPostID(postId uint) (dtos.PostResponse, error) {
	post, err := p.postRepo.GetPostById(postId)
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

// UpdatePost implements PostService.
func (p *postService) UpdatePost(req dtos.UpdatePost) (dtos.PostResponse, error) {
	if err := validator.New().Struct(&req); err != nil {
		return dtos.PostResponse{}, err
	}

	exsistingPost, err := p.postRepo.GetPostById(req.ID)
	if err != nil {
		return dtos.PostResponse{}, err
	}

	if exsistingPost.UserID != req.UserID {
		return dtos.PostResponse{}, errors.New("unauthorized to update this post")
	}

	if req.Content != nil {
		exsistingPost.Content = *req.Content
	}
	if req.ImageURL != nil {
		exsistingPost.ImageURL = req.ImageURL
	}

	post, err := p.postRepo.UpdatePost(exsistingPost)
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
