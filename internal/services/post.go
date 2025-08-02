package services

import (
	"context"
	"time"

	"github.com/hanzala211/CRUD/internal/api/models"
	"github.com/hanzala211/CRUD/internal/repo"
)

type PostService struct {
	postRepo *repo.PostRepo
}

func NewPostService(postRepo *repo.PostRepo) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, post *models.Post) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.postRepo.CreatePost(ctx, post)
}

func (s *PostService) GetPostByID(ctx context.Context, post *models.Post, postID string) error {
	post.ID = postID
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.postRepo.GetPostByID(ctx, post)
}
