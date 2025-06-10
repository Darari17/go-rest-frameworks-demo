package usecases

import (
	"errors"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/models"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/repositories"
)

type PostUsecase struct {
	postRepo repositories.PostRepo
}

func NewPostUsecase(postRepo repositories.PostRepo) *PostUsecase {
	return &PostUsecase{
		postRepo: postRepo,
	}
}

func (p *PostUsecase) CreatePost(req dtos.CreatePost) (dtos.PostResponse, error) {
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

func (p *PostUsecase) DeletePost(postId uint) error {
	return p.postRepo.DeletePost(postId)
}

func (p *PostUsecase) FindPostsByUserId(userId uint) ([]dtos.PostResponse, error) {
	posts, err := p.postRepo.FindPostsByUserId(userId)
	if err != nil {
		return nil, err
	}

	responses := []dtos.PostResponse{}
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

func (p *PostUsecase) FindPostByPostId(postId uint) (dtos.PostResponse, error) {
	post, err := p.postRepo.FindPostByPostId(postId)
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

func (p *PostUsecase) UpdatePost(req dtos.UpdatePost) (dtos.PostResponse, error) {
	existingPost, err := p.postRepo.FindPostByPostId(req.ID)
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

	update, err := p.postRepo.UpdatePost(existingPost)
	if err != nil {
		return dtos.PostResponse{}, err
	}

	response := dtos.PostResponse{
		ID:        update.ID,
		UserID:    update.UserID,
		Content:   update.Content,
		ImageURL:  update.ImageURL,
		CreatedAt: update.CreatedAt,
		UpdatedAt: update.UpdatedAt,
	}

	return response, nil
}
