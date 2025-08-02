package services

import (
	"context"
	"time"

	"github.com/hanzala211/CRUD/internal/api/models"
	"github.com/hanzala211/CRUD/internal/repo"
)

type CommentService struct {
	commentRepo *repo.CommentRepo
}

func NewCommentService(commentRepo *repo.CommentRepo) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (s *CommentService) AddComment(ctx context.Context, comment *models.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.commentRepo.AddComment(ctx, comment)
}

func (s *CommentService) GetPostComments(ctx context.Context, postID string) ([]*models.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.commentRepo.GetPostComments(ctx, postID)
}
