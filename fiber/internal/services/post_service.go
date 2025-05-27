package services

import (
	"errors"

	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/models"
	"github.com/Darari17/go-rest-frameworks-demo/fiber/internal/repositories"
	"github.com/go-playground/validator/v10"
)

type postService struct {
	postRepo repositories.PostRepo
}

type PostService interface {
	CreatePostService(req dtos.CreatePost) (dtos.PostResponse, error)
	FindByPostIDService(postId uint) (dtos.PostResponse, error)
	FindAllPostService() ([]dtos.PostResponse, error)
	UpdatePostService(req dtos.UpdatePost) (dtos.PostResponse, error)
	DeletePostService(postId uint) error
}

func NewPostService(postRepo repositories.PostRepo) PostService {
	return &postService{
		postRepo: postRepo,
	}
}

// CreatePost implements PostService.
func (p *postService) CreatePostService(req dtos.CreatePost) (dtos.PostResponse, error) {
	if err := validator.New().Struct(&req); err != nil {
		return dtos.PostResponse{}, err
	}

	post, err := p.postRepo.CreatePostRepo(&models.Post{
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
func (p *postService) DeletePostService(postId uint) error {
	return p.postRepo.DeletePostRepo(postId)
}

// FindAllPost implements PostService.
func (p *postService) FindAllPostService() ([]dtos.PostResponse, error) {
	posts, err := p.postRepo.GetPostsRepo()
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
func (p *postService) FindByPostIDService(postId uint) (dtos.PostResponse, error) {
	post, err := p.postRepo.GetPostByIdRepo(postId)
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
func (p *postService) UpdatePostService(req dtos.UpdatePost) (dtos.PostResponse, error) {
	if err := validator.New().Struct(&req); err != nil {
		return dtos.PostResponse{}, err
	}

	exsistingPost, err := p.postRepo.GetPostByIdRepo(req.ID)
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

	post, err := p.postRepo.UpdatePostRepo(exsistingPost)
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
