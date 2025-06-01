package services

import (
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/jwt"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/repositories"
)

type PostService struct {
	postRepo   repositories.PostRepo
	jwtHandler jwt.JWTHandler
}

func NewPostService(postRepo repositories.PostRepo, jwtHandler jwt.JWTHandler) *PostService {
	return &PostService{
		postRepo:   postRepo,
		jwtHandler: jwtHandler,
	}
}

// func (ps *PostService) CreatePost()
// func (ps *PostService) DeletePost()
// func (ps *PostService) UpdatePost()
// func (ps *PostService) GetPostByPostID()
// func (ps *PostService) GetPostsByUserID()
