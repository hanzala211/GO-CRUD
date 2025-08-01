package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/hanzala211/CRUD/internal/api/models"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	_, err := r.db.ModelContext(ctx, user).Insert()
	return err
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *models.User, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	user.ID = userId
	defer cancel()
	// _, err := r.db.ModelContext(ctx, user).WherePK().Update()
	_, err := r.db.ModelContext(ctx, user).Where("id = ?", userId).Update()
	return err
}

func (r *UserRepo) DeleteUser(ctx context.Context, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	_, err := r.db.ModelContext(ctx, &models.User{}).Where("id = ?", userId).Delete()
	return err
}

func (r *UserRepo) GetUsers(ctx context.Context) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	users := []*models.User{}
	defer cancel()
	err := r.db.ModelContext(ctx, &users).Select()
	fmt.Println(err)
	return users, err
}

func (r *UserRepo) GetUserByID(ctx context.Context, user *models.User, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	err := r.db.ModelContext(ctx, user).Where("id = ?", userId).Select()
	return err
}
