package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/hanzala211/CRUD/internal/api/models"
)

type CommentRepo struct {
	db *pg.DB
}

func NewCommentRepo(db *pg.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

func (r *CommentRepo) AddComment(ctx context.Context, comment *models.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	_, err := r.db.ModelContext(ctx, comment).Insert()
	return err
}

func (r *CommentRepo) GetPostComments(ctx context.Context, postId string) ([]*models.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	comments := []*models.Comment{}
	err := r.db.ModelContext(ctx, &comments).Relation("User").Where("post_id = ?", postId).Select()
	fmt.Print(err)
	return comments, err
}

func (r *CommentRepo) TestFunc(ctx context.Context, comment any) error {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	_, err := r.db.ModelContext(ctx, comment).Insert()
	return err
}
