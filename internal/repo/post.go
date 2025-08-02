package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/hanzala211/CRUD/internal/api/models"
)

type PostRepo struct {
	db *pg.DB
}

func NewPostRepo(db *pg.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) CreatePost(ctx context.Context, post *models.Post) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	_, err := r.db.ModelContext(ctx, post).Insert()
	fmt.Printf("Error: %v", err)
	return err
}

func (r *PostRepo) GetPostByID(ctx context.Context, post *models.Post) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	err := r.db.ModelContext(ctx, post).Relation("User").WherePK().Select()
	return err
}
